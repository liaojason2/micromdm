package queue

import (
	"net/url"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/micromdm/micromdm/pkg/httputil"
)

func NewHTTPClient(instance, token string, logger log.Logger, opts ...httptransport.ClientOption) (Service, error) {
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}

	var getDeviceCommandEndpoint endpoint.Endpoint
	{
		getDeviceCommandEndpoint = httptransport.NewClient(
			"POST",
			httputil.CopyURL(u, "/v1/queue"),
			httputil.EncodeRequestWithToken(token, httptransport.EncodeJSONRequest),
			decodeGetDeviceCommandResponse,
			opts...,
		).Endpoint()
	}

	return Endpoints{
		GetDeviceCommandEndpoint: getDeviceCommandEndpoint,
	}, nil
}
