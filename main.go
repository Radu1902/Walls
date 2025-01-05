package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/wall", getWall)

	http.ListenAndServe(":8080", nil)
}
