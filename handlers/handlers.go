package handlers

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

const MysqlDataSource = "root:password@tcp(127.0.0.1:3306)/catalog"

type CarMark struct {
	IdCarMark string `json:"id_car_mark"`
	Name      string `json:"name"`
}

type CarModification struct {
	IdCarModification string `json:"id_car_modification"`
	Name              string `json:"name"`
}

type CarGeneration struct {
	IdCarGeneration string `json:"id_car_generation"`
	Name            string `json:"name"`
	YearBegin       string `json:"year_begin"`
	YearEnd         string `json:"year_end"`
}

type CarSerie struct {
	IdCarSerie       string            `json:"id_car_serie"`
	Name             string            `json:"name"`
	CarModifications []CarModification `json:"car_modifications"`
}

type CarModel struct {
	IdCarModel     string          `json:"id_car_model"`
	Name           string          `json:"name"`
	CarGenerations []CarGeneration `json:"car_generations"`
}

func HandleCarSerie(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", MysqlDataSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var idKey = req.URL.Query().Get("key")

	results, err := db.Query(
		"SELECT id_car_serie, name "+
			"FROM car_serie WHERE id_car_generation = ?", idKey)
	if err != nil {
		panic(err.Error())
	}

	var resp []CarSerie
	for results.Next() {
		var carserie CarSerie
		err = results.Scan(&carserie.IdCarSerie, &carserie.Name)
		if err != nil {
			panic(err.Error())
		}

		results2, err := db.Query(
			"SELECT id_car_modification, name "+
				"FROM car_modification WHERE id_car_serie = ?", carserie.IdCarSerie)
		if err != nil {
			panic(err.Error())
		}

		for results2.Next() {
			var carmodif CarModification
			err = results2.Scan(&carmodif.IdCarModification, &carmodif.Name)
			if err != nil {
				panic(err.Error())
			}

			carserie.CarModifications = append(carserie.CarModifications, carmodif)
		}

		resp = append(resp, carserie)
	}

	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(resp)
	w.Write(response)
}

func HandleCarModels(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", MysqlDataSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var idCarMark = req.URL.Query().Get("key")

	results, err := db.Query(
		"SELECT id_car_model, name "+
			"FROM car_model WHERE id_car_mark = ?", idCarMark)
	if err != nil {
		panic(err.Error())
	}

	var resp []CarModel
	for results.Next() {
		var carmodel CarModel
		err = results.Scan(&carmodel.IdCarModel, &carmodel.Name)
		if err != nil {
			panic(err.Error())
		}

		results2, err := db.Query(
			"SELECT id_car_generation, name, year_begin, year_end "+
				"FROM car_generation WHERE id_car_model = ? ORDER BY year_begin", carmodel.IdCarModel)
		if err != nil {
			panic(err.Error())
		}

		for results2.Next() {
			var cargen CarGeneration
			err = results2.Scan(&cargen.IdCarGeneration, &cargen.Name, &cargen.YearBegin, &cargen.YearEnd)
			if err != nil {
				panic(err.Error())
			}
			carmodel.CarGenerations = append(carmodel.CarGenerations, cargen)
		}

		resp = append(resp, carmodel)
	}

	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(resp)
	w.Write(response)
}

func HandleCarMakes(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", MysqlDataSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT id_car_mark, name FROM car_mark WHERE id_car_type = 1")
	if err != nil {
		panic(err.Error())
	}

	var resp []CarMark
	for results.Next() {
		var carmark CarMark
		err = results.Scan(&carmark.IdCarMark, &carmark.Name)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, carmark)
	}

	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(resp)
	w.Write(response)
}
