package transports

import (
	"context"
	"github.com/autobaza/auto_catalog/endpoints"
	"github.com/autobaza/auto_catalog/protos"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	listCarTypes               gt.Handler
	listCarMarks               gt.Handler
	listCarModels              gt.Handler
	listCarGenerations         gt.Handler
	listCarSeriesByModel       gt.Handler
	listCarSeriesByGeneration  gt.Handler
	listCarModifications       gt.Handler
	listCarEquipments          gt.Handler
	listCarCharacteristicValue gt.Handler
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
			decodeCarRequest,
			encodeListCarMarksResponse,
		),
		listCarModels: gt.NewServer(
			endpoints.ListCarModels,
			decodeCarRequest,
			encodeListCarModelsResponse,
		),
		listCarGenerations: gt.NewServer(
			endpoints.ListCarGenerations,
			decodeCarRequest,
			encodeListCarGenerationsResponse,
		),
		listCarSeriesByModel: gt.NewServer(
			endpoints.ListCarSeriesByModel,
			decodeCarRequest,
			encodeListCarSeriesResponse,
		),
		listCarSeriesByGeneration: gt.NewServer(
			endpoints.ListCarSeriesByGeneration,
			decodeCarRequest,
			encodeListCarSeriesResponse,
		),
		listCarModifications: gt.NewServer(
			endpoints.ListCarModifications,
			decodeCarRequest,
			encodeListCarModificationsResponse,
		),
		listCarEquipments: gt.NewServer(
			endpoints.ListCarEquipments,
			decodeCarRequest,
			encodeListCarEquipmentsResponse,
		),
		listCarCharacteristicValue: gt.NewServer(
			endpoints.ListCarCharacteristicValues,
			decodeCarRequest,
			encodeListCarCharacteristicValueResponse,
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

func (g gRPCServer) ListCarGenerations(ctx context.Context, req *catalog.CarRequest) (*catalog.ListCarGenerationResponse, error) {
	_, resp, err := g.listCarGenerations.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*catalog.ListCarGenerationResponse), nil
}

func (g gRPCServer) ListCarSeriesByModel(ctx context.Context, request *catalog.CarRequest) (*catalog.ListCarSeriesResponse, error) {
	_, resp, err := g.listCarSeriesByModel.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*catalog.ListCarSeriesResponse), nil
}

func (g gRPCServer) ListCarSeriesByGeneration(ctx context.Context, request *catalog.CarRequest) (*catalog.ListCarSeriesResponse, error) {
	_, resp, err := g.listCarSeriesByGeneration.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*catalog.ListCarSeriesResponse), nil
}

func (g gRPCServer) ListCarModifications(ctx context.Context, request *catalog.CarRequest) (*catalog.ListCarModificationsResponse, error) {
	_, resp, err := g.listCarModifications.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*catalog.ListCarModificationsResponse), nil
}

func (g gRPCServer) ListCarEquipments(ctx context.Context, request *catalog.CarRequest) (*catalog.ListCarEquipmentsResponse, error) {
	_, resp, err := g.listCarEquipments.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*catalog.ListCarEquipmentsResponse), nil
}

func (g gRPCServer) ListCarCharacteristicValue(ctx context.Context, request *catalog.CarRequest) (*catalog.ListCarCharacteristicValueResponse, error) {
	_, resp, err := g.listCarCharacteristicValue.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*catalog.ListCarCharacteristicValueResponse), nil
}

func decodeListCarTypesRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeListCarTypesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ListCarTypesResponse)
	return &catalog.ListCarTypesResponse{CarTypes: resp.CarTypes}, nil
}

func decodeCarRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*catalog.CarRequest)
	return endpoints.CarRequest{ID: req.Id}, nil
}

func encodeListCarMarksResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ListCarMarksResponse)
	return &catalog.ListCarMarkResponse{CarMarks: resp.CarMarks}, nil
}

func encodeListCarModelsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ListCarModelsResponse)
	return &catalog.ListCarModelResponse{CarModels: resp.CarModels}, nil
}

func encodeListCarGenerationsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ListCarGenerationsResponse)
	return &catalog.ListCarGenerationResponse{CarGenerations: resp.CarGenerations}, nil
}

func encodeListCarSeriesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ListCarSeriesResponse)
	return &catalog.ListCarSeriesResponse{CarSeries: resp.CarSeries}, nil
}

func encodeListCarModificationsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ListCarModificationsResponse)
	return &catalog.ListCarModificationsResponse{CarModifications: resp.CarModifications}, nil
}

func encodeListCarEquipmentsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ListCarEquipmentsResponse)
	return &catalog.ListCarEquipmentsResponse{CarEquipments: resp.CarEquipments}, nil
}

func encodeListCarCharacteristicValueResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.ListCarCharacteristicValues)
	return &catalog.ListCarCharacteristicValueResponse{CarCharacteristics: resp.CarCharacteristicValues}, nil
}
