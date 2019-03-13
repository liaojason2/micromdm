package queue

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/micromdm/micromdm/pkg/httputil"
)

func (svc *DeviceCommandService) GetDeviceCommand(ctx context.Context, udid string) (DeviceCommand, error) {
	return svc.store.DeviceCommand(udid)
}

func decodeGetDeviceCommandRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req string
	err := httputil.DecodeJSONRequest(r, &req)
	return req, err
}

func decodeGetDeviceCommandResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var resp DeviceCommand
	err := httputil.DecodeJSONResponse(r, &resp)
	return resp, err
}

func MakeGetDeviceCommandEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		udid := request.(string)
		return GetDeviceCommand(ctx, udid)
	}
}

func (e Endpoints) GetDeviceCommand(ctx context.Context, udid string) (DeviceCommand, error) {
	response, err := e.GetDeviceCommandEndpoint(ctx, udid)
	if err != nil {
		return nil, err
	}
	return response, err
}
