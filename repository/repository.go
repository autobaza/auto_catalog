package repository

import (
	"database/sql"
	catalog "github.com/autobaza/auto_catalog/protos"
	_ "github.com/go-sql-driver/mysql"
)

type Repository interface {
	GetCarTypes() []*catalog.CarType
	GetCarMarks(string) []*catalog.CarMark
	GetCarModels(string) []*catalog.CarModel
	GetCarGenerations(string) []*catalog.CarGeneration
	GetCarSeriesByModel(string) []*catalog.CarSerie
	GetCarSeriesByGeneration(string) []*catalog.CarSerie
	GetCarModifications(string) []*catalog.CarModification
	GetCarEquipments(string) []*catalog.CarEquipment
	GetCarCharacteristicValue(string) []*catalog.CarCharacteristicValue
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

func (r *repository) GetCarModels(id string) []*catalog.CarModel {
	results, err := r.db.Query("SELECT id_car_model, id_car_mark, name, coalesce(name_rus, '') FROM car_model WHERE id_car_mark = ?", id)
	if err != nil {
		panic(err.Error())
	}
	var resp []*catalog.CarModel
	for results.Next() {
		var carmodel struct {
			Id      string `json:"id_car_model"`
			MarkId  string `json:"id_car_mark"`
			Name    string `json:"name"`
			NameRus string `json:"name_rus"`
		}
		err = results.Scan(&carmodel.Id, &carmodel.MarkId, &carmodel.Name, &carmodel.NameRus)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, &catalog.CarModel{Id: carmodel.Id, Name: carmodel.Name, MarkId: carmodel.MarkId, NameRus: carmodel.NameRus})
	}
	return resp
}

func (r *repository) GetCarGenerations(id string) []*catalog.CarGeneration {
	results, err := r.db.Query(
		"SELECT id_car_generation, id_car_model, name, coalesce(year_begin, ''), coalesce(year_end, '') "+
			"FROM car_generation WHERE id_car_model = ? ORDER BY year_begin", id)
	if err != nil {
		panic(err.Error())
	}
	var resp []*catalog.CarGeneration
	for results.Next() {
		var cargen struct {
			Id        string `json:"id_car_generation"`
			ModelId   string `json:"id_car_model"`
			Name      string `json:"name"`
			YearBegin string `json:"year_begin"`
			YearEnd   string `json:"year_end"`
		}
		err = results.Scan(&cargen.Id, &cargen.ModelId, &cargen.Name, &cargen.YearBegin, &cargen.YearEnd)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, &catalog.CarGeneration{
			Id:        cargen.Id,
			ModelId:   cargen.ModelId,
			Name:      cargen.Name,
			YearBegin: cargen.YearBegin,
			YearEnd:   cargen.YearEnd,
		})
	}
	return resp
}

func (r *repository) GetCarSeriesByModel(id string) []*catalog.CarSerie {
	return r.GetCarSeries(id, "model")
}

func (r *repository) GetCarSeriesByGeneration(id string) []*catalog.CarSerie {
	return r.GetCarSeries(id, "generation")
}

func (r *repository) GetCarSeries(id string, method string) []*catalog.CarSerie {
	var query = "SELECT id_car_serie, id_car_model, name, id_car_generation FROM car_serie WHERE"
	if method == "model" {
		query += " id_car_model = ?"
	} else {
		query += " id_car_generation = ?"
	}
	results, err := r.db.Query(query, id)
	if err != nil {
		panic(err.Error())
	}
	var resp []*catalog.CarSerie
	for results.Next() {
		var carserie struct {
			Id           string `json:"id_car_serie"`
			ModelId      string `json:"id_car_model"`
			Name         string `json:"name"`
			GenerationId string `json:"id_car_generation"`
		}
		err = results.Scan(&carserie.Id, &carserie.ModelId, &carserie.Name, &carserie.GenerationId)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, &catalog.CarSerie{
			Id:           carserie.Id,
			ModelId:      carserie.ModelId,
			Name:         carserie.Name,
			GenerationId: carserie.GenerationId,
		})
	}
	return resp
}

func (r *repository) GetCarModifications(id string) []*catalog.CarModification {
	results, err := r.db.Query(
		"SELECT id_car_modification, id_car_serie, id_car_model, name, coalesce(start_production_year, ''), coalesce(end_production_year, '') "+
			"FROM car_modification WHERE id_car_serie = ?", id)
	if err != nil {
		panic(err.Error())
	}
	var resp []*catalog.CarModification
	for results.Next() {
		var carserie struct {
			Id            string `json:"id_car_modification"`
			SerieId       string `json:"id_car_serie"`
			ModelId       string `json:"id_car_model"`
			Name          string `json:"name"`
			StartProdYear string `json:"start_production_year"`
			EndProdYear   string `json:"end_production_year"`
		}
		err = results.Scan(&carserie.Id, &carserie.SerieId, &carserie.ModelId, &carserie.Name, &carserie.StartProdYear, &carserie.EndProdYear)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, &catalog.CarModification{
			Id:                  carserie.Id,
			SerieId:             carserie.SerieId,
			ModelId:             carserie.ModelId,
			Name:                carserie.Name,
			StartProductionYear: carserie.StartProdYear,
			EndProductionYear:   carserie.EndProdYear,
		})
	}
	return resp
}

func (r *repository) GetCarEquipments(id string) []*catalog.CarEquipment {
	results, err := r.db.Query(
		"SELECT id_car_equipment, id_car_modification, name, coalesce(price_min, ''), coalesce(year, '') "+
			"FROM car_equipment WHERE id_car_modification = ?", id)
	if err != nil {
		panic(err.Error())
	}
	var resp []*catalog.CarEquipment
	for results.Next() {
		var carequip struct {
			Id       string `json:"id_car_equipment"`
			ModifId  string `json:"id_car_modification"`
			Name     string `json:"name"`
			PriceMin string `json:"price_min"`
			Year     string `json:"year"`
		}
		err = results.Scan(&carequip.Id, &carequip.ModifId, &carequip.Name, &carequip.PriceMin, &carequip.Year)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, &catalog.CarEquipment{
			Id:             carequip.Id,
			ModificationId: carequip.ModifId,
			Name:           carequip.Name,
			PriceMin:       carequip.PriceMin,
			Year:           carequip.Year,
		})
	}
	return resp
}

func (r *repository) GetCarCharacteristicValue(id string) []*catalog.CarCharacteristicValue {
	//TODO implement me
	panic("implement me")
}
