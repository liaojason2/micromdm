package queue

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/micromdm/micromdm/pkg/httputil"
)

type getDeviceCommandRequest struct {
	UDID string `json:"udid"`
}

func (svc *DeviceCommandService) GetDeviceCommand(ctx context.Context, udid string) (DeviceCommand, error) {
	device, err := svc.store.DeviceCommand(udid)
	return *device, err
}

func decodeGetDeviceCommandRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getDeviceCommandRequest
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
		return svc.GetDeviceCommand(ctx, "CEB23C56-E653-5B9F-8DC1-85120CA33CD5")
	}
}

func (e Endpoints) GetDeviceCommand(ctx context.Context, udid string) (DeviceCommand, error) {
	response, err := e.GetDeviceCommandEndpoint(ctx, udid)
	return response.(DeviceCommand), err
}
