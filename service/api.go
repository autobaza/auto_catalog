package service

import (
	"context"
	"database/sql"
	catalog "github.com/autobaza/auto_catalog/protos"
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

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
		db:     connectToDatabase(),
	}
}

func connectToDatabase() *sql.DB {
	var dsn = "root:password@tcp(127.0.0.1:3306)/catalog"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type service struct {
	db     *sql.DB
	logger log.Logger
}

type ListCarTypes struct {
	list []catalog.CarType
}

func (s *service) ListCarTypes(ctx context.Context) []*catalog.CarType {

	var resp []*catalog.CarType
	results, err := s.db.Query("SELECT id_car_type, name FROM car_type ORDER BY id_car_type")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var cartype struct {
			Id   string `json:"id_car_type"`
			Name string `json:"name"`
		}
		err = results.Scan(&cartype.Id, &cartype.Name)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, &catalog.CarType{Id: cartype.Id, Name: cartype.Name})
	}
	return resp
}

func (s *service) ListCarMarks(ctx context.Context, id string) []*catalog.CarMark {

	results, err := s.db.Query("SELECT id_car_mark, name, id_car_type, coalesce(name_rus, '') FROM car_mark WHERE id_car_type = ?", id)
	if err != nil {
		panic(err.Error())
	}

	type CarMark struct {
		Id      string `json:"id_car_mark"`
		Name    string `json:"name"`
		TypeId  string `json:"id_car_type"`
		NameRus string `json:"name_rus"`
	}
	var resp []*catalog.CarMark
	for results.Next() {
		var carmark CarMark
		err = results.Scan(&carmark.Id, &carmark.Name, &carmark.TypeId, &carmark.NameRus)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, &catalog.CarMark{Id: carmark.Id, Name: carmark.Name, TypeId: carmark.TypeId, NameRus: carmark.NameRus})
	}
	return resp
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
