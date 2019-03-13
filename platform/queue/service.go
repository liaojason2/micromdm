package queue

import (
	"context"
)

type Service interface {
	GetDeviceCommand(ctx context.Context, udid string) (DeviceCommand, error)
}

type Store interface {
	DeviceCommand(udid string) (DeviceCommand, error)
}

type DeviceCommandService struct {
	store Store
}

func New(store Store) *DeviceCommandService {
	return &DeviceCommandService{store: store}
}
