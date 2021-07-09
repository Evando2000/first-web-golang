package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("GO + My SQL")

	username_db := "root"
	password_db := "postgresql"
	host_db := "127.0.0.1"
	port_db := "3306"
	name_db := "gomysqlproj1"

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username_db, password_db, host_db, port_db, name_db),
	)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// ========= Insert new User =========
	// new_user := "'Echo'"
	// insert, err := db.Query("INSERT INTO users VALUES(" + new_user + ")")

	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()

	// fmt.Println(new_user + " registered")

	// ========= See all Users =========
	results, err := db.Query("SELECT name FROM users")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var user User
		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.Name)
	}

}
