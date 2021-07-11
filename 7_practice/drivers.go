package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Driver struct {
	gorm.Model
	Name    string
	License string
	Cars    []Car
}

var (
	drivers = []Driver{
		{Name: "Jimmy Johnson", License: "ABC123"},
		{Name: "Howard Hills", License: "XYZ789"},
		{Name: "Craig Colbin", License: "DEF333"},
	}
)

func GetDrivers(w http.ResponseWriter, r *http.Request) {
	var drivers []Driver
	db.Find(&drivers)
	json.NewEncoder(w).Encode(&drivers)
}

func GetDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var driver Driver
	var cars []Car
	db.First(&driver, params["id"])
	db.Model(&driver).Related(&cars)
	driver.Cars = cars
	json.NewEncoder(w).Encode(&driver)
}

func CreateDriver(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	name := r.FormValue("name")
	license := r.FormValue("license")
	newDriver := &Driver{Name: name, License: license}
	db.Create(newDriver)
	json.NewEncoder(w).Encode(&newDriver)
}
