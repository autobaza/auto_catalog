package endpoints

import (
	"context"
	catalog "github.com/autobaza/auto_catalog/protos"
	"github.com/autobaza/auto_catalog/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ListCarTypes endpoint.Endpoint
	ListCarMarks endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		ListCarTypes: makeListCarTypesEndpoint(s),
		ListCarMarks: makeListCarMarksEndpoint(s),
	}
}

func makeListCarTypesEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return ListCarTypesResponse{CarTypes: s.ListCarTypes(ctx)}, nil
	}
}

func makeListCarMarksEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CarRequest)
		return ListCarMarksResponse{CarMarks: s.ListCarMarks(ctx, req.Id)}, nil
	}
}

type CarRequest struct {
	Id string
}

type ListCarTypesResponse struct {
	CarTypes []*catalog.CarType
}

type ListCarMarksResponse struct {
	CarMarks []*catalog.CarMark
}
