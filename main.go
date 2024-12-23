package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

var users []User = []User{{"ionut", "qwer", []byte("salut")}, {"mircea", "asdf", []byte("buna")}}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body any
	decoder.Decode(&body)
	fmt.Println()
	fmt.Println(r.Body)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	randVar := rand.Uint32() % 3
	json.NewEncoder(w).Encode(randVar)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	usr := users[1]
	fmt.Println(usr)
	json.NewEncoder(w).Encode(usr)
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/auth", authHandler)
	http.HandleFunc("/user", getUser)
	http.HandleFunc("/register", registerHandler)

	http.ListenAndServe(":8080", nil)
}
