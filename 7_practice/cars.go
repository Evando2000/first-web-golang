package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Car struct {
	gorm.Model
	Year      int
	Make      string
	ModelName string
	DriverID  int
}

var (
	cars = []Car{
		{Year: 2000, Make: "Toyota", ModelName: "Tundra", DriverID: 1},
		{Year: 2001, Make: "Honda", ModelName: "Accord", DriverID: 1},
		{Year: 2002, Make: "Nissan", ModelName: "Sentra", DriverID: 2},
		{Year: 2003, Make: "Ford", ModelName: "F-150", DriverID: 3},
	}
)

func GetCars(w http.ResponseWriter, r *http.Request) {
	var cars []Car
	db.Find(&cars)
	json.NewEncoder(w).Encode(&cars)
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var car Car
	db.First(&car, params["id"])
	json.NewEncoder(w).Encode(&car)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var car Car
	db.First(&car, params["id"])
	db.Delete(&car)
	var cars []Car
	db.Find(&cars)
	json.NewEncoder(w).Encode(&cars)
}
