package queue

import (
	"context"
)

type Service interface {
	GetDeviceCommand(ctx context.Context, udid string) (DeviceCommand, error)
}

type QueueStore interface {
	GetDeviceCommand(ctx context.Context, udid string) (DeviceCommand, error)
}

type QueueService struct {
	store QueueStore
}

func New(store QueueStore) *QueueService {
	return &QueueService{store: store}
}
