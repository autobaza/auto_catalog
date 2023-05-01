package service

import (
	"context"
	"database/sql"
	catalog "github.com/autobaza/auto_catalog/protos"
	"github.com/go-kit/kit/log"
	_ "github.com/go-sql-driver/mysql"
)

const MysqlDataSource = "root:password@tcp(127.0.0.1:3306)/catalog"

type Service interface {
	ListCarTypes(ctx context.Context) []*catalog.CarType
	//GetCarMarks([]int)
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
	}
}

type service struct {
	db     *sql.DB
	logger log.Logger
}

type ListCarTypes struct {
	list []catalog.CarType
}

func (s *service) ListCarTypes(ctx context.Context) []*catalog.CarType {
	db, err := sql.Open("mysql", MysqlDataSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT id_car_type, name FROM car_type ORDER BY id_car_type")
	if err != nil {
		panic(err.Error())
	}

	type CarType struct {
		Id   string `json:"id_car_type"`
		Name string `json:"name"`
	}
	var resp []*catalog.CarType
	for results.Next() {
		var cartype CarType
		err = results.Scan(&cartype.Id, &cartype.Name)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, &catalog.CarType{Id: cartype.Id, Name: cartype.Name})
	}
	return resp
}

//func (s *service) GetCarMarks(carTypeIds []int) {
//
//}
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
