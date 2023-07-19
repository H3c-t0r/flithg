package implementations

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"time"

	dataInterfaces "github.com/flyteorg/flyteadmin/pkg/data/interfaces"
	"github.com/flyteorg/flyteadmin/pkg/manager/impl/util"
	repositoryInterfaces "github.com/flyteorg/flyteadmin/pkg/repositories/interfaces"
	"github.com/flyteorg/flyteadmin/pkg/repositories/transformers"
	runtimeInterfaces "github.com/flyteorg/flyteadmin/pkg/runtime/interfaces"
	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/event"
	"github.com/flyteorg/flytestdlib/storage"

	"github.com/golang/protobuf/jsonpb"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/flyteorg/flyteadmin/pkg/async/notifications/implementations"
	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin"

	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/flyteorg/flyteadmin/pkg/async/cloudevent/interfaces"

	"github.com/flyteorg/flytestdlib/logger"
	"github.com/flyteorg/flytestdlib/promutils"
	"github.com/golang/protobuf/proto"
)

const (
	cloudEventSource     = "https://github.com/flyteorg/flyteadmin"
	cloudEventTypePrefix = "com.flyte.resource"
	jsonSchemaURLKey     = "jsonschemaurl"
	jsonSchemaURL        = "https://github.com/flyteorg/flyteidl/blob/v0.24.14/jsonschema/workflow_execution.json"
)

// Publisher This event publisher acts to asynchronously publish workflow execution events.
type Publisher struct {
	sender        interfaces.Sender
	systemMetrics implementations.EventPublisherSystemMetrics
	events        sets.String
}

func (p *Publisher) Publish(ctx context.Context, notificationType string, msg proto.Message) error {
	if !p.shouldPublishEvent(notificationType) {
		return nil
	}
	p.systemMetrics.PublishTotal.Inc()
	logger.Debugf(ctx, "Publishing the following message [%+v]", msg)

	var executionID string
	var phase string
	var eventTime time.Time

	switch msgType := msg.(type) {
	case *admin.WorkflowExecutionEventRequest:
		e := msgType.Event
		executionID = e.ExecutionId.String()
		phase = e.Phase.String()
		eventTime = e.OccurredAt.AsTime()
	case *admin.TaskExecutionEventRequest:
		e := msgType.Event
		executionID = e.TaskId.String()
		phase = e.Phase.String()
		eventTime = e.OccurredAt.AsTime()
	case *admin.NodeExecutionEventRequest:
		e := msgType.Event
		executionID = msgType.Event.Id.String()
		phase = e.Phase.String()
		eventTime = e.OccurredAt.AsTime()
	default:
		return fmt.Errorf("unsupported event types [%+v]", reflect.TypeOf(msg))
	}

	event := cloudevents.NewEvent()
	// CloudEvent specification: https://github.com/cloudevents/spec/blob/v1.0/spec.md#required-attributes
	event.SetType(fmt.Sprintf("%v.%v", cloudEventTypePrefix, notificationType))
	event.SetSource(cloudEventSource)
	event.SetID(fmt.Sprintf("%v.%v", executionID, phase))
	event.SetTime(eventTime)
	event.SetExtension(jsonSchemaURLKey, jsonSchemaURL)

	// Explicitly jsonpb marshal the proto. Otherwise, event.SetData will use json.Marshal which doesn't work well
	// with proto oneof fields.
	marshaler := jsonpb.Marshaler{}
	buf := bytes.NewBuffer([]byte{})
	err := marshaler.Marshal(buf, msg)
	if err != nil {
		p.systemMetrics.PublishError.Inc()
		logger.Errorf(ctx, "Failed to jsonpb marshal [%v] with error: %v", msg, err)
		return fmt.Errorf("failed to jsonpb marshal [%v] with error: %w", msg, err)
	}

	if err := event.SetData(cloudevents.ApplicationJSON, buf.Bytes()); err != nil {
		p.systemMetrics.PublishError.Inc()
		logger.Errorf(ctx, "Failed to encode message [%v] with error: %v", msg, err)
		return err
	}

	if err := p.sender.Send(ctx, notificationType, event); err != nil {
		p.systemMetrics.PublishError.Inc()
		logger.Errorf(ctx, "Failed to send message [%v] with error: %v", msg, err)
		return err
	}
	p.systemMetrics.PublishSuccess.Inc()
	return nil
}

func (p *Publisher) shouldPublishEvent(notificationType string) bool {
	return p.events.Has(notificationType)
}

type CloudEventWrappedPublisher struct {
	db               repositoryInterfaces.Repository
	sender           interfaces.Sender
	systemMetrics    implementations.EventPublisherSystemMetrics
	storageClient    *storage.DataStore
	urlData          dataInterfaces.RemoteURLInterface
	remoteDataConfig runtimeInterfaces.RemoteDataConfig
}

func (c *CloudEventWrappedPublisher) TransformWorkflowExecutionEvent(ctx context.Context, rawEvent *event.WorkflowExecutionEvent) (*event.CloudEventWorkflowExecution, error) {

	// Basic error checking
	if rawEvent == nil {
		return nil, fmt.Errorf("nothing to publish, WorkflowExecution event is nil")
	}
	if rawEvent.ExecutionId == nil {
		logger.Warningf(ctx, "nil execution id in event [%+v]", rawEvent)
		return nil, fmt.Errorf("nil execution id in event [%+v]", rawEvent)
	}

	// TODO: Make this one call to the DB instead of two.
	executionModel, err := c.db.ExecutionRepo().Get(ctx, repositoryInterfaces.Identifier{
		Project: rawEvent.ExecutionId.Project,
		Domain:  rawEvent.ExecutionId.Domain,
		Name:    rawEvent.ExecutionId.Name,
	})
	ex, err := transformers.FromExecutionModel(ctx, executionModel, nil)
	if ex.Closure.WorkflowId == nil {
		logger.Warningf(ctx, "workflow id is nil for execution [%+v]", ex)
		return nil, fmt.Errorf("workflow id is nil for execution [%+v]", ex)
	}
	workflowModel, err := c.db.WorkflowRepo().Get(ctx, repositoryInterfaces.Identifier{
		Project: ex.Closure.WorkflowId.Project,
		Domain:  ex.Closure.WorkflowId.Domain,
		Name:    ex.Closure.WorkflowId.Name,
		Version: ex.Closure.WorkflowId.Version,
	})
	var workflowInterface core.TypedInterface
	if workflowModel.TypedInterface != nil && len(workflowModel.TypedInterface) > 0 {
		err = proto.Unmarshal(workflowModel.TypedInterface, &workflowInterface)
		if err != nil {
			return nil, fmt.Errorf(
				"artifact eventing - failed to unmarshal TypedInterface for workflow [%+v] with err: %v",
				workflowModel.ID, err)
		}
	}

	var outputs *core.LiteralMap
	if rawEvent.GetOutputData() != nil {
		fmt.Printf("remove this - Got output data")
		outputs = rawEvent.GetOutputData()
	} else if len(rawEvent.GetOutputUri()) > 0 {
		fmt.Printf("remove this - Got output URI")
		// GetInputs actually fetches the data, even though this is an output
		outputs, _, err = util.GetInputs(ctx, c.urlData, &c.remoteDataConfig, c.storageClient, rawEvent.GetOutputUri())
		if err != nil {
			// TODO: metric this
			logger.Warningf(ctx, "Error fetching output literal map %v", rawEvent)
			return nil, err
		}
	}

	if outputs == nil {
		// todo: remove after testing
		logger.Debugf(ctx, "Output data was nil for %v", rawEvent)
	}

	return &event.CloudEventWorkflowExecution{
		RawEvent:        rawEvent,
		OutputData:      outputs,
		OutputInterface: &workflowInterface,
	}, nil
}

func (c *CloudEventWrappedPublisher) TransformNodeExecutionEvent(ctx context.Context, rawEvent *event.NodeExecutionEvent) (*event.CloudEventNodeExecution, error) {
	return &event.CloudEventNodeExecution{
		RawEvent: rawEvent,
	}, nil
}

func (c *CloudEventWrappedPublisher) TransformTaskExecutionEvent(ctx context.Context, rawEvent *event.TaskExecutionEvent) (*event.CloudEventTaskExecution, error) {
	if rawEvent == nil {
		return nil, fmt.Errorf("nothing to publish, TaskExecution event is nil")
	}

	taskModel, err := c.db.TaskRepo().Get(ctx, repositoryInterfaces.Identifier{
		Project: rawEvent.TaskId.Project,
		Domain:  rawEvent.TaskId.Domain,
		Name:    rawEvent.TaskId.Name,
		Version: rawEvent.TaskId.Version,
	})
	if err != nil {
		// TODO: metric this
		logger.Debugf(ctx, "Failed to get task with task id [%+v] with err %v", rawEvent.TaskId, err)
		return nil, err
	}
	task, err := transformers.FromTaskModel(taskModel)

	var outputs *core.LiteralMap
	if rawEvent.GetOutputData() != nil {
		fmt.Printf("remove this - task Got output data")
		outputs = rawEvent.GetOutputData()
	} else if len(rawEvent.GetOutputUri()) > 0 {
		fmt.Printf("remove this - task Got output URI")
		// GetInputs actually fetches the data, even though this is an output
		outputs, _, err = util.GetInputs(ctx, c.urlData, &c.remoteDataConfig,
			c.storageClient, rawEvent.GetOutputUri())
		if err != nil {
			fmt.Printf("Error fetching output literal map %v", rawEvent)
			return nil, err
		}
	}
	if outputs == nil {
		// todo: remove
		fmt.Printf("No output data found for task execution %v\n", rawEvent)
	}

	return &event.CloudEventTaskExecution{
		RawEvent:        rawEvent,
		OutputData:      outputs,
		OutputInterface: task.Closure.CompiledTask.Template.Interface,
	}, nil
}

func (c *CloudEventWrappedPublisher) Publish(ctx context.Context, notificationType string, msg proto.Message) error {
	c.systemMetrics.PublishTotal.Inc()
	logger.Debugf(ctx, "Publishing the following message [%+v]", msg)

	var err error
	var executionID string
	var phase string
	var eventTime time.Time
	var finalMsg proto.Message

	switch msgType := msg.(type) {
	case *admin.WorkflowExecutionEventRequest:
		e := msgType.Event
		executionID = e.ExecutionId.String()
		phase = e.Phase.String()
		eventTime = e.OccurredAt.AsTime()

		finalMsg, err = c.TransformWorkflowExecutionEvent(ctx, e)
		if err != nil {
			logger.Errorf(ctx, "Failed to transform workflow execution event with error: %v", err)
			return err
		}

	case *admin.TaskExecutionEventRequest:
		e := msgType.Event
		executionID = e.TaskId.String()
		phase = e.Phase.String()
		eventTime = e.OccurredAt.AsTime()
		finalMsg, err = c.TransformTaskExecutionEvent(ctx, e)
	case *admin.NodeExecutionEventRequest:
		e := msgType.Event
		executionID = msgType.Event.Id.String()
		phase = e.Phase.String()
		eventTime = e.OccurredAt.AsTime()
		finalMsg, err = c.TransformNodeExecutionEvent(ctx, e)
	default:
		return fmt.Errorf("unsupported event types [%+v]", reflect.TypeOf(msg))
	}

	// Explicitly jsonpb marshal the proto. Otherwise, event.SetData will use json.Marshal which doesn't work well
	// with proto oneof fields.
	marshaler := jsonpb.Marshaler{}
	buf := bytes.NewBuffer([]byte{})
	err = marshaler.Marshal(buf, finalMsg)
	if err != nil {
		c.systemMetrics.PublishError.Inc()
		logger.Errorf(ctx, "Failed to jsonpb marshal [%v] with error: %v", msg, err)
		return fmt.Errorf("failed to jsonpb marshal [%v] with error: %w", msg, err)
	}

	cloudEvt := cloudevents.NewEvent()
	// CloudEvent specification: https://github.com/cloudevents/spec/blob/v1.0/spec.md#required-attributes
	cloudEvt.SetType(fmt.Sprintf("%v.%v", cloudEventTypePrefix, notificationType))
	cloudEvt.SetSource(cloudEventSource)
	cloudEvt.SetID(fmt.Sprintf("%v.%v", executionID, phase))
	cloudEvt.SetTime(eventTime)
	cloudEvt.SetExtension(jsonSchemaURLKey, jsonSchemaURL)

	if err := cloudEvt.SetData(cloudevents.ApplicationJSON, buf.Bytes()); err != nil {
		c.systemMetrics.PublishError.Inc()
		logger.Errorf(ctx, "Failed to encode message [%v] with error: %v", msg, err)
		return err
	}

	if err := c.sender.Send(ctx, notificationType, cloudEvt); err != nil {
		c.systemMetrics.PublishError.Inc()
		logger.Errorf(ctx, "Failed to send message [%v] with error: %v", msg, err)
		return err
	}
	c.systemMetrics.PublishSuccess.Inc()
	return nil
}

func NewCloudEventsPublisher(sender interfaces.Sender, scope promutils.Scope, eventTypes []string) interfaces.Publisher {
	eventSet := sets.NewString()

	for _, eventType := range eventTypes {
		if eventType == implementations.AllTypes || eventType == implementations.AllTypesShort {
			for _, e := range implementations.SupportedEvents {
				eventSet = eventSet.Insert(e)
			}
			break
		}
		if e, found := implementations.SupportedEvents[eventType]; found {
			eventSet = eventSet.Insert(e)
		} else {
			panic(fmt.Errorf("unsupported event type [%s] in the config", eventType))
		}
	}

	return &Publisher{
		sender:        sender,
		systemMetrics: implementations.NewEventPublisherSystemMetrics(scope.NewSubScope("cloudevents_publisher")),
		events:        eventSet,
	}
}

func NewCloudEventsWrappedPublisher(
	db repositoryInterfaces.Repository, sender interfaces.Sender, scope promutils.Scope, storageClient *storage.DataStore, urlData dataInterfaces.RemoteURLInterface, remoteDataConfig runtimeInterfaces.RemoteDataConfig) interfaces.Publisher {

	return &CloudEventWrappedPublisher{
		db:               db,
		sender:           sender,
		systemMetrics:    implementations.NewEventPublisherSystemMetrics(scope.NewSubScope("cloudevents_publisher")),
		storageClient:    storageClient,
		urlData:          urlData,
		remoteDataConfig: remoteDataConfig,
	}
}
