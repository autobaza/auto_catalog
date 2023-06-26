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
			ID   string `json:"id_car_type"`
			Name string `json:"name"`
		}
		err = results.Scan(&cartype.ID, &cartype.Name)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, &catalog.CarType{Id: cartype.ID, Name: cartype.Name})
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
			ID      string `json:"id_car_mark"`
			Name    string `json:"name"`
			TypeID  string `json:"id_car_type"`
			NameRus string `json:"name_rus"`
		}
		err = results.Scan(&carmark.ID, &carmark.Name, &carmark.TypeID, &carmark.NameRus)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, &catalog.CarMark{Id: carmark.ID, Name: carmark.Name, TypeId: carmark.TypeID, NameRus: carmark.NameRus})
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
			ID      string `json:"id_car_model"`
			MarkID  string `json:"id_car_mark"`
			Name    string `json:"name"`
			NameRus string `json:"name_rus"`
		}
		err = results.Scan(&carmodel.ID, &carmodel.MarkID, &carmodel.Name, &carmodel.NameRus)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, &catalog.CarModel{Id: carmodel.ID, Name: carmodel.Name, MarkId: carmodel.MarkID, NameRus: carmodel.NameRus})
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
			ID        string `json:"id_car_generation"`
			ModelID   string `json:"id_car_model"`
			Name      string `json:"name"`
			YearBegin string `json:"year_begin"`
			YearEnd   string `json:"year_end"`
		}
		err = results.Scan(&cargen.ID, &cargen.ModelID, &cargen.Name, &cargen.YearBegin, &cargen.YearEnd)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, &catalog.CarGeneration{
			Id:        cargen.ID,
			ModelId:   cargen.ModelID,
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
			ID           string `json:"id_car_serie"`
			ModelID      string `json:"id_car_model"`
			Name         string `json:"name"`
			GenerationID string `json:"id_car_generation"`
		}
		err = results.Scan(&carserie.ID, &carserie.ModelID, &carserie.Name, &carserie.GenerationID)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, &catalog.CarSerie{
			Id:           carserie.ID,
			ModelId:      carserie.ModelID,
			Name:         carserie.Name,
			GenerationId: carserie.GenerationID,
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
			ID            string `json:"id_car_modification"`
			SerieID       string `json:"id_car_serie"`
			ModelID       string `json:"id_car_model"`
			Name          string `json:"name"`
			StartProdYear string `json:"start_production_year"`
			EndProdYear   string `json:"end_production_year"`
		}
		err = results.Scan(&carserie.ID, &carserie.SerieID, &carserie.ModelID, &carserie.Name, &carserie.StartProdYear, &carserie.EndProdYear)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, &catalog.CarModification{
			Id:                  carserie.ID,
			SerieId:             carserie.SerieID,
			ModelId:             carserie.ModelID,
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
			ID       string `json:"id_car_equipment"`
			ModifID  string `json:"id_car_modification"`
			Name     string `json:"name"`
			PriceMin string `json:"price_min"`
			Year     string `json:"year"`
		}
		err = results.Scan(&carequip.ID, &carequip.ModifID, &carequip.Name, &carequip.PriceMin, &carequip.Year)
		if err != nil {
			panic(err.Error())
		}
		resp = append(resp, &catalog.CarEquipment{
			Id:             carequip.ID,
			ModificationId: carequip.ModifID,
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
