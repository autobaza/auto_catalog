package endpoints

import (
	"context"
	catalog "github.com/autobaza/auto_catalog/protos"
	"github.com/autobaza/auto_catalog/service"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	ListCarTypes                endpoint.Endpoint
	ListCarMarks                endpoint.Endpoint
	ListCarModels               endpoint.Endpoint
	ListCarGenerations          endpoint.Endpoint
	ListCarSeriesByModel        endpoint.Endpoint
	ListCarSeriesByGeneration   endpoint.Endpoint
	ListCarModifications        endpoint.Endpoint
	ListCarEquipments           endpoint.Endpoint
	ListCarCharacteristicValues endpoint.Endpoint
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		ListCarTypes:                makeListCarTypesEndpoint(s),
		ListCarMarks:                makeListCarMarksEndpoint(s),
		ListCarModels:               makeListCarModelsEndpoint(s),
		ListCarGenerations:          makeListCarGenerationsEndpoint(s),
		ListCarSeriesByModel:        makeListCarSeriesByModelEndpoint(s),
		ListCarSeriesByGeneration:   makeListCarSeriesByGenerationEndpoint(s),
		ListCarModifications:        makeListCarModificationsEndpoint(s),
		ListCarEquipments:           makeListCarEquipmentsEndpoint(s),
		ListCarCharacteristicValues: makeListCarCharacteristicValuesEndpoint(s),
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
		return ListCarMarksResponse{CarMarks: s.ListCarMarks(ctx, req.ID)}, nil
	}
}

func makeListCarModelsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CarRequest)
		return ListCarModelsResponse{CarModels: s.ListCarModels(ctx, req.ID)}, nil
	}
}

func makeListCarGenerationsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CarRequest)
		return ListCarGenerationsResponse{CarGenerations: s.ListCarGenerations(ctx, req.ID)}, nil
	}
}

func makeListCarSeriesByModelEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CarRequest)
		return ListCarSeriesResponse{CarSeries: s.ListCarSeriesByModel(ctx, req.ID)}, nil
	}
}

func makeListCarSeriesByGenerationEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CarRequest)
		return ListCarSeriesResponse{CarSeries: s.ListCarSeriesByGeneration(ctx, req.ID)}, nil
	}
}

func makeListCarModificationsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CarRequest)
		return ListCarModificationsResponse{CarModifications: s.ListCarModifications(ctx, req.ID)}, nil
	}
}

func makeListCarEquipmentsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CarRequest)
		return ListCarEquipmentsResponse{CarEquipments: s.ListCarEquipments(ctx, req.ID)}, nil
	}
}

func makeListCarCharacteristicValuesEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CarRequest)
		return ListCarCharacteristicValues{CarCharacteristicValues: s.ListCarCharacteristicValue(ctx, req.ID)}, nil
	}
}

type CarRequest struct {
	ID string
}

type ListCarTypesResponse struct {
	CarTypes []*catalog.CarType
}

type ListCarMarksResponse struct {
	CarMarks []*catalog.CarMark
}

type ListCarModelsResponse struct {
	CarModels []*catalog.CarModel
}

type ListCarGenerationsResponse struct {
	CarGenerations []*catalog.CarGeneration
}

type ListCarSeriesResponse struct {
	CarSeries []*catalog.CarSerie
}

type ListCarModificationsResponse struct {
	CarModifications []*catalog.CarModification
}

type ListCarEquipmentsResponse struct {
	CarEquipments []*catalog.CarEquipment
}

type ListCarCharacteristicValues struct {
	CarCharacteristicValues []*catalog.CarCharacteristicValue
}
