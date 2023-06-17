package service

import (
	"context"
	catalog "github.com/autobaza/auto_catalog/protos"
	"github.com/autobaza/auto_catalog/repository"
	"github.com/go-kit/kit/log"
	_ "github.com/go-sql-driver/mysql"
)

type Service interface {
	ListCarTypes(context.Context) []*catalog.CarType
	ListCarMarks(context.Context, string) []*catalog.CarMark
	ListCarModels(context.Context, string) []*catalog.CarModel
	ListCarGenerations(context.Context, string) []*catalog.CarGeneration
	ListCarSeriesByModel(context.Context, string) []*catalog.CarSerie
	ListCarSeriesByGeneration(context.Context, string) []*catalog.CarSerie
	ListCarModifications(context.Context, string) []*catalog.CarModification
	ListCarEquipments(context.Context, string) []*catalog.CarEquipment
	ListCarCharacteristicValue(context.Context, string) []*catalog.CarCharacteristicValue
}

func NewService(logger log.Logger, repo repository.Repository) Service {
	return &service{
		logger: logger,
		repo:   repo,
	}
}

type service struct {
	repo   repository.Repository
	logger log.Logger
}

type ListCarTypes struct {
	list []catalog.CarType
}

func (s *service) ListCarTypes(ctx context.Context) []*catalog.CarType {
	return s.repo.GetCarTypes()
}

func (s *service) ListCarMarks(ctx context.Context, id string) []*catalog.CarMark {
	return s.repo.GetCarMarks(id)
}

func (s *service) ListCarModels(ctx context.Context, id string) []*catalog.CarModel {
	return s.repo.GetCarModels(id)
}

func (s *service) ListCarGenerations(ctx context.Context, id string) []*catalog.CarGeneration {
	return s.repo.GetCarGenerations(id)
}

func (s *service) ListCarSeriesByModel(ctx context.Context, id string) []*catalog.CarSerie {
	return s.repo.GetCarSeriesByModel(id)
}

func (s *service) ListCarSeriesByGeneration(ctx context.Context, id string) []*catalog.CarSerie {
	return s.repo.GetCarSeriesByGeneration(id)
}

func (s *service) ListCarModifications(ctx context.Context, id string) []*catalog.CarModification {
	return s.repo.GetCarModifications(id)
}

func (s *service) ListCarEquipments(ctx context.Context, id string) []*catalog.CarEquipment {
	return s.repo.GetCarEquipments(id)
}

func (s *service) ListCarCharacteristicValue(ctx context.Context, id string) []*catalog.CarCharacteristicValue {
	return s.repo.GetCarCharacteristicValue(id)
}
