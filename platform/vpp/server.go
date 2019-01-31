package vpp

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/micromdm/micromdm/pkg/httputil"
)

type Endpoints struct {
	GetAssetsSrvEndpoint                 endpoint.Endpoint
	GetLicensesSrvEndpoint               endpoint.Endpoint
	ManageVPPLicensesByAdamIdSrvEndpoint endpoint.Endpoint
	GetServiceConfigSrvEndpoint          endpoint.Endpoint
}

func MakeServerEndpoints(s Service, outer endpoint.Middleware, others ...endpoint.Middleware) Endpoints {
	return Endpoints{
		GetAssetsSrvEndpoint:                 endpoint.Chain(outer, others...)(MakeGetAssetsSrvEndpoint(s)),
		GetLicensesSrvEndpoint:               endpoint.Chain(outer, others...)(MakeGetLicensesSrvEndpoint(s)),
		ManageVPPLicensesByAdamIdSrvEndpoint: endpoint.Chain(outer, others...)(MakeManageVPPLicensesByAdamIdSrvEndpoint(s)),
		GetServiceConfigSrvEndpoint:          endpoint.Chain(outer, others...)(MakeGetServiceConfigSrvEndpoint(s)),
	}
}

func RegisterHTTPHandlers(r *mux.Router, e Endpoints, options ...httptransport.ServerOption) {
	// GET		/v1/vpp/assets	        list all vpp assets
	// POST		/v1/vpp/assets		      list all vpp assets with options
	// GET		/v1/vpp/licenses		    list all vpp licenses
	// POST		/v1/vpp/licenses		    list all vpp licenses with options
	// PUT		/v1/vpp/licenses				manage vpp licenses
	// GET		/v1/vpp/serviceconfig		get vpp service config information
	// POST		/v1/vpp/serviceconfig		get vpp service config information for specific sToken

	r.Methods("GET").Path("/v1/vpp/assets").Handler(httptransport.NewServer(
		e.GetAssetsSrvEndpoint,
		decodeGetAssetsSrvRequest,
		httputil.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/v1/vpp/assets").Handler(httptransport.NewServer(
		e.GetAssetsSrvEndpoint,
		decodeGetAssetsSrvRequest,
		httputil.EncodeJSONResponse,
		options...,
	))

	r.Methods("GET").Path("/v1/vpp/licenses").Handler(httptransport.NewServer(
		e.GetLicensesSrvEndpoint,
		decodeGetLicensesSrvRequest,
		httputil.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/v1/vpp/licenses").Handler(httptransport.NewServer(
		e.GetLicensesSrvEndpoint,
		decodeGetLicensesSrvRequest,
		httputil.EncodeJSONResponse,
		options...,
	))

	r.Methods("PUT").Path("/v1/vpp/licenses").Handler(httptransport.NewServer(
		e.ManageVPPLicensesByAdamIdSrvEndpoint,
		decodeManageVPPLicensesByAdamIdSrvRequest,
		httputil.EncodeJSONResponse,
		options...,
	))

	r.Methods("GET").Path("/v1/vpp/serviceconfig").Handler(httptransport.NewServer(
		e.GetServiceConfigSrvEndpoint,
		decodeGetServiceConfigSrvRequest,
		httputil.EncodeJSONResponse,
		options...,
	))

	r.Methods("POST").Path("/v1/vpp/serviceconfig").Handler(httptransport.NewServer(
		e.GetServiceConfigSrvEndpoint,
		decodeGetServiceConfigSrvRequest,
		httputil.EncodeJSONResponse,
		options...,
	))
}
