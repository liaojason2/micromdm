package queue

import (
	"context"

	"github.com/micromdm/micromdm/platform/pubsub"
)

type Service interface {
	GetDeviceCommand(ctx context.Context, udid string) (DeviceCommand, error)
}

type QueueService interface {
	publisher pubsub.Publisher
	GetDeviceCommand(string) (DeviceCommand, error)
}

func New(pub pubsub.Publisher) *QueueService {
	svc := QueueService{
		publisher: pub,
	}
	return &svc
}
