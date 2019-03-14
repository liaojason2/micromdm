package queue

import (
	"context"

	"github.com/micromdm/micromdm/platform/pubsub"
)

type Service interface {
	GetDeviceCommand(ctx context.Context, udid string) (DeviceCommand, error)
}

type QueueService struct {
	publisher pubsub.Publisher
}

func New(pub pubsub.Publisher) (*QueueService, error) {
	svc := QueueService{
		publisher: pub,
	}
	return &svc, nil
}
