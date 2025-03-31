package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/glebarez/go-sqlite"
)

type Database struct {
	db *sql.DB
}

func main() {
	db, err := sql.Open("sqlite", "./users.db")

	if err != nil {
		fmt.Println(err)
		return
	} else {
		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users ( username TEXT PRIMARY KEY, password TEXT NOT NULL, wall BLOB );`)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("database instantiated successfully")
		}
	}
	database := Database{db: db}
	err = database.init()
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/register", database.registerHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/wall/", getWall)
	http.HandleFunc("/update", updateWallHandler)

	http.ListenAndServe(":8080", nil)

	defer db.Close()
}
