package queue

import (
	"context"
)

type Service interface {
	GetDeviceCommand(ctx context.Context, udid string) (DeviceCommand, error)
}

type QueueService struct {
	store Store
}

func New(store Store) *QueueService {
	return &QueueService{store: store}
}
