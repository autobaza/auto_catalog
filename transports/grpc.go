package transports

import (
	"context"
	"github.com/autobaza/auto_catalog/endpoints"
	"github.com/autobaza/auto_catalog/protos"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	listCarTypes gt.Handler
}

func (g gRPCServer) ListCarTypes(ctx context.Context, req *catalog.Empty) (*catalog.ListCarTypesResponse, error) {
	_, resp, err := g.listCarTypes.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*catalog.ListCarTypesResponse), nil
}

func NewGRPCServer(endpoints endpoints.Endpoints) catalog.AutoCatalogServiceServer {
	return &gRPCServer{
		listCarTypes: gt.NewServer(
			endpoints.ListCarTypes,
			decodeListCarTypesRequest,
			encodeListCarTypesResponse,
		),
	}
}

func decodeListCarTypesRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeListCarTypesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ListCarTypesResponse)
	return &catalog.ListCarTypesResponse{CarTypes: resp.CarTypes}, nil
}
