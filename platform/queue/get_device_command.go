package queue

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	"github.com/micromdm/micromdm/pkg/httputil"
)

func (svc *QueueService) GetDeviceCommand(ctx context.Context, udid string) (DeviceCommand, error) {
	response, err := svc.store.DeviceCommand(udid)
	return *response, err
}

type deviceCommandRequest struct {
	UDID string `json:"udid"`
}

func decodeGetDeviceCommandRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req deviceCommandRequest
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
		req := request.(deviceCommandRequest)
		details, err := svc.GetDeviceCommand(ctx, req.UDID)
		return details, err
	}
}

func (e Endpoints) GetDeviceCommand(ctx context.Context, udid string) (DeviceCommand, error) {
	request := deviceCommandRequest{UDID: udid}
	response, err := e.GetDeviceCommandEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(DeviceCommand), nil
}
