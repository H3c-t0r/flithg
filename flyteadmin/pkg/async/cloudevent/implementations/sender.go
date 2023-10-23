package implementations

import (
	"context"
	"fmt"
	pbcloudevents "github.com/cloudevents/sdk-go/binding/format/protobuf/v2"

	"github.com/NYTimes/gizmo/pubsub"
	"github.com/Shopify/sarama"
	"github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/flyteorg/flyte/flytestdlib/logger"
)

type Receiver = string

const (
	Kafka Receiver = "Kafka"
)

// PubSubSender Implementation of Sender
type PubSubSender struct {
	Pub pubsub.Publisher
}

func (s *PubSubSender) Send(ctx context.Context, notificationType string, event cloudevents.Event) error {
	// gatepr: investigate why the previous statement didn't work.
	//  perhaps only because of redis.
	eventByte, err := pbcloudevents.Protobuf.Marshal(&event)
	//eventByte, err := json.Marshal(&event)
	if err != nil {
		logger.Errorf(ctx, "Failed to marshal cloudevent with error: %v", err)
		return err
	}
	if err := s.Pub.PublishRaw(ctx, notificationType, eventByte); err != nil {
		logger.Errorf(ctx, "Failed to publish a message with key [%s] and message [%s] and error: %v", notificationType, event.String(), err)
		return err
	}

	return nil
}

// KafkaSender Implementation of Sender
type KafkaSender struct {
	Client cloudevents.Client
}

func (s *KafkaSender) Send(ctx context.Context, notificationType string, event cloudevents.Event) error {
	if result := s.Client.Send(
		// Set the producer message key
		kafka_sarama.WithMessageKey(ctx, sarama.StringEncoder(event.ID())),
		event,
	); cloudevents.IsUndelivered(result) {
		return fmt.Errorf("failed to send cloud event: %v", result)
	}
	return nil
}
