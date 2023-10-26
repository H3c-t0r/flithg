package processor

import (
	"bytes"
	"context"
	"fmt"
	pbcloudevents "github.com/cloudevents/sdk-go/binding/format/protobuf/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	flyteEvents "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/event"
	"github.com/flyteorg/flyte/flytestdlib/logger"
	"github.com/flyteorg/flyte/flytestdlib/sandbox_utils"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"time"
)

type SandboxCloudEventsReceiver struct {
	subChan <-chan sandbox_utils.SandboxMessage
	Handler EventsHandlerInterface
}

func (p *SandboxCloudEventsReceiver) StartProcessing(ctx context.Context) {
	for {
		logger.Warningf(context.Background(), "Starting SandBox notifications processor")
		err := p.run(ctx)
		if err != nil {
			logger.Errorf(context.Background(), "error with running processor err: [%v], sleeping and restarting", err)
			time.Sleep(1000 * 1000 * 1000 * 1)
		}
		break
	}
	logger.Warning(context.Background(), "Sandbox cloud event processor has stopped because context cancelled")
}

func (p *SandboxCloudEventsReceiver) handleMessage(ctx context.Context, sandboxMsg sandbox_utils.SandboxMessage) error {
	ce := &event.Event{}
	err := pbcloudevents.Protobuf.Unmarshal(sandboxMsg.Raw, ce)
	if err != nil {
		logger.Errorf(context.Background(), "error with unmarshalling message [%v]", err)
		return err
	}
	logger.Debugf(ctx, "Cloud event received message [%+v]", ce)
	// ce data should be a jsonpb Marshaled proto message, one of
	// - event.CloudEventTaskExecution
	// - event.CloudEventNodeExecution
	// - event.CloudEventWorkflowExecution
	// - event.CloudEventExecutionStart
	ceData := bytes.NewReader(ce.Data())
	unmarshaler := jsonpb.Unmarshaler{}

	// Use the type to determine which proto message to unmarshal to.
	var flyteEvent proto.Message
	if ce.Type() == "com.flyte.resource.cloudevents.TaskExecution" {
		flyteEvent = &flyteEvents.CloudEventTaskExecution{}
		err = unmarshaler.Unmarshal(ceData, flyteEvent)
	} else if ce.Type() == "com.flyte.resource.cloudevents.WorkflowExecution" {
		flyteEvent = &flyteEvents.CloudEventWorkflowExecution{}
		err = unmarshaler.Unmarshal(ceData, flyteEvent)
	} else if ce.Type() == "com.flyte.resource.cloudevents.NodeExecution" {
		flyteEvent = &flyteEvents.CloudEventNodeExecution{}
		err = unmarshaler.Unmarshal(ceData, flyteEvent)
	} else if ce.Type() == "com.flyte.resource.cloudevents.ExecutionStart" {
		flyteEvent = &flyteEvents.CloudEventExecutionStart{}
		err = unmarshaler.Unmarshal(ceData, flyteEvent)
	} else {
		logger.Warningf(ctx, "Ignoring cloud event type [%s]", ce.Type())
		return nil
	}
	if err != nil {
		logger.Errorf(ctx, "error unmarshalling message on topic [%s] [%v]", sandboxMsg.Topic, err)
		return err
	}

	err = p.Handler.HandleEvent(ctx, ce, flyteEvent)
	if err != nil {
		logger.Errorf(context.Background(), "error handling event on topic [%s] [%v]", sandboxMsg.Topic, err)
		return err
	}
	return nil
}

func (p *SandboxCloudEventsReceiver) run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			logger.Warning(context.Background(), "Context cancelled, stopping processing.")
			return nil

		case sandboxMsg := <-p.subChan:
			fmt.Println("HERE222!!!!!!!!!!-receive")
			logger.Debugf(ctx, "received message [%v]", sandboxMsg)
			if sandboxMsg.Raw != nil {
				err := p.handleMessage(ctx, sandboxMsg)
				if err != nil {
					// Assuming that handle message will return a fair number of errors
					// add metric
					logger.Infof(ctx, "error processing sandbox cloud event [%v] with err [%v]", sandboxMsg, err)
				}
				fmt.Println("HERE222!!!!!!!!!!-success")
			} else {
				logger.Infof(ctx, "sandbox receiver ignoring message [%v]", sandboxMsg)
			}
		}
	}
}

func (p *SandboxCloudEventsReceiver) StopProcessing() error {
	logger.Warning(context.Background(), "StopProcessing called on SandboxCloudEventsReceiver")
	return nil
}

func NewSandboxCloudEventProcessor(eventsHandler EventsHandlerInterface) *SandboxCloudEventsReceiver {
	return &SandboxCloudEventsReceiver{
		Handler: eventsHandler,
		subChan: sandbox_utils.MsgChan,
	}
}
