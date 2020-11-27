package domain

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

import "google.golang.org/api/pubsub/v1"


//create a new MessageBus that forwards all messages to PubSub
func newPubSubMessageBus(topic string) (MessageBus, error) {
	var ctx = context.Background()
	service, err := pubsub.NewService(ctx)
	publisher := pubsub.NewProjectsTopicsService(service)

	if err != nil {
		logrus.Errorf("unable to create Pubsub message bus: %v", err)
		return nil, err
	}
	return &pubsubMessageBus{ctx, publisher, topic}, nil
}

type pubsubMessageBus struct {
	ctx context.Context
	service *pubsub.ProjectsTopicsService
	topic string
}

func (p *pubsubMessageBus) Send(id string, data []byte) error {
	fmt.Printf("data: %s\r\n", string(data))
	encoded := base64.StdEncoding.EncodeToString(data)

	call := p.service.Publish(p.topic, &pubsub.PublishRequest{
		Messages:       []*pubsub.PubsubMessage{{
			Attributes:      nil,
			Data:            encoded,
			MessageId:       id,
			PublishTime:     time.Now().Format(time.RFC3339),
			ForceSendFields: nil,
			NullFields:      nil,
				}},
		ForceSendFields: nil,
		NullFields:      nil,
	})

	_ , err := call.Do()
	if err != nil {
		logrus.Errorf("unable to publish event: %v", err)
	}

	return nil
}
