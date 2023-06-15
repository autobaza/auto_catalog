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
	//GetCarModels(int)
	//GetCarGenerations(int)
	//GetCarSeriesByModel(int)
	//getCarSeriesByGeneration(int)
	//GetCarModifications(int)
	//GetCarEquipment(int)
	//GetCarCharacteristicValue(int)
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

//
//func (s *service) GetCarModels(carMarkId int) {
//
//}
//
//func (s *service) GetCarGenerations(carModelId int) {
//
//}
//func (s *service) GetCarSeriesByModel(carModelId int) {
//
//}
//func (s *service) getCarSeriesByGeneration(carGenerationId int) {
//
//}
//func (s *service) GetCarModifications(carSerieId int) {
//
//}
//func (s *service) GetCarEquipment(carModificationId int) {
//
//}
//func (s *service) GetCarCharacteristicValue(carModificationId int) {
//
//}
