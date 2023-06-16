package transports

import (
	"context"
	"github.com/autobaza/auto_catalog/endpoints"
	"github.com/autobaza/auto_catalog/protos"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	listCarTypes  gt.Handler
	listCarMarks  gt.Handler
	listCarModels gt.Handler
}

func NewGRPCServer(endpoints endpoints.Endpoints) catalog.AutoCatalogServiceServer {
	return &gRPCServer{
		listCarTypes: gt.NewServer(
			endpoints.ListCarTypes,
			decodeListCarTypesRequest,
			encodeListCarTypesResponse,
		),
		listCarMarks: gt.NewServer(
			endpoints.ListCarMarks,
			decodeListCarMarksRequest,
			encodeListCarMarksResponse,
		),
		listCarModels: gt.NewServer(
			endpoints.ListCarModels,
			decodeListCarModelsRequest,
			encodeListCarModelsResponse,
		),
	}
}

func (g gRPCServer) ListCarTypes(ctx context.Context, req *catalog.Empty) (*catalog.ListCarTypesResponse, error) {
	_, resp, err := g.listCarTypes.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*catalog.ListCarTypesResponse), nil
}

func (g gRPCServer) ListCarMarks(ctx context.Context, req *catalog.CarRequest) (*catalog.ListCarMarkResponse, error) {
	_, resp, err := g.listCarMarks.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*catalog.ListCarMarkResponse), nil
}

func (g gRPCServer) ListCarModels(ctx context.Context, req *catalog.CarRequest) (*catalog.ListCarModelResponse, error) {
	_, resp, err := g.listCarModels.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*catalog.ListCarModelResponse), nil
}

func decodeListCarTypesRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeListCarTypesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ListCarTypesResponse)
	return &catalog.ListCarTypesResponse{CarTypes: resp.CarTypes}, nil
}

func decodeListCarMarksRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*catalog.CarRequest)
	return endpoints.CarRequest{Id: req.Id}, nil
}

func encodeListCarMarksResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ListCarMarksResponse)
	return &catalog.ListCarMarkResponse{CarMarks: resp.CarMarks}, nil
}

func decodeListCarModelsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*catalog.CarRequest)
	return endpoints.CarRequest{Id: req.Id}, nil
}

func encodeListCarModelsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ListCarModelsResponse)
	return &catalog.ListCarModelResponse{CarModels: resp.CarModels}, nil
}
