package queue

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/micromdm/micromdm/pkg/httputil"
)

type Endpoints struct {
	GetDeviceCommandEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service, outer endpoint.Middleware, others ...endpoint.Middleware) Endpoints {
	return Endpoints{
		GetDeviceCommandEndpoint: endpoint.Chain(outer, others...)(MakeGetDeviceCommandEndpoint(s)),
	}
}

func RegisterHTTPHandlers(r *mux.Router, e Endpoints, options ...httptransport.ServerOption) {
	// POST     /v1/queue		get the command queue for a device managed by the server

	r.Methods("POST").Path("/v1/queue").Handler(httptransport.NewServer(
		e.GetDeviceCommandEndpoint,
		decodeGetDeviceCommandRequest,
		httputil.EncodeJSONResponse,
		options...,
	))

}
