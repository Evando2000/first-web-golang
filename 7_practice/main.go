package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
)

var db *gorm.DB
var err error

func urlHandler() {
	router := mux.NewRouter()
	router.HandleFunc("/cars", GetCars).Methods("GET")
	router.HandleFunc("/drivers", GetDrivers).Methods("GET")
	router.HandleFunc("/drivers", CreateDriver).Methods("POST")
	router.HandleFunc("/cars/{id}", GetCar).Methods("GET")
	router.HandleFunc("/drivers/{id}", GetDriver).Methods("GET")
	router.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func migration() {
	if err != nil {
		panic("failed to connect database")
	}
	db.Exec("CREATE SCHEMA IF NOT EXISTS gopostgre")
	db.Exec("DROP TABLE IF EXISTS gopostgre.drivers")
	db.Exec("DROP TABLE IF EXISTS gopostgre.cars")

	// set schema here.
	gorm.DefaultTableNameHandler = func(db *gorm.DB, tableName string) string {
		return "gopostgre." + tableName
	}

	// db.DropTableIfExists(&Driver{})
	// db.DropTableIfExists(&Car{})
	db.AutoMigrate(&Driver{})
	db.AutoMigrate(&Car{})

	for index := range cars {
		db.Create(&cars[index])
	}

	for index := range drivers {
		db.Create(&drivers[index])
	}
}

func main() {
	host_db := "localhost"
	port_db := "5432"
	username_db := "postgres"
	name_db := "postgres"
	sslmode := "disable"
	password_db := "postgresql"

	db, err = gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
			host_db, port_db, username_db, name_db, sslmode, password_db),
	)

	migration()
	urlHandler()
	defer db.Close()
}
