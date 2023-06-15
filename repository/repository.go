package repository

import (
	"database/sql"
	catalog "github.com/autobaza/auto_catalog/protos"
)

type Repository interface {
	GetCarTypes() []*catalog.CarType
	GetCarMarks(string) []*catalog.CarMark
}

type repository struct {
	db *sql.DB
}

func NewRepository() Repository {
	db := connectDatabase()
	return &repository{
		db: db,
	}
}

func connectDatabase() *sql.DB {
	var dsn = "root:password@tcp(127.0.0.1:3306)/catalog"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func (r *repository) GetCarTypes() []*catalog.CarType {
	var resp []*catalog.CarType
	results, err := r.db.Query("SELECT id_car_type, name FROM car_type ORDER BY id_car_type")
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

func (r *repository) GetCarMarks(id string) []*catalog.CarMark {
	results, err := r.db.Query("SELECT id_car_mark, name, id_car_type, coalesce(name_rus, '') FROM car_mark WHERE id_car_type = ?", id)
	if err != nil {
		panic(err.Error())
	}
	var resp []*catalog.CarMark
	for results.Next() {
		var carmark struct {
			Id      string `json:"id_car_mark"`
			Name    string `json:"name"`
			TypeId  string `json:"id_car_type"`
			NameRus string `json:"name_rus"`
		}
		err = results.Scan(&carmark.Id, &carmark.Name, &carmark.TypeId, &carmark.NameRus)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, &catalog.CarMark{Id: carmark.Id, Name: carmark.Name, TypeId: carmark.TypeId, NameRus: carmark.NameRus})
	}
	return resp
}
