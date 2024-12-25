package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var users []User = []User{{"ionut", "qwer", []byte("salut")}, {"mircea", "asdf", []byte("buna")}}
var tokens map[string][]byte = make(map[string][]byte)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body string
	decoder.Decode(&body)
	creds := strings.Fields(body)
	newUser := User{creds[0], creds[1], []byte("")}
	users = append(users, newUser)
	// to do: db sync
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body string
	decoder.Decode(&body)
	creds := strings.Fields(body)
	fmt.Println(creds[0], creds[1])
	for index := range users {
		if users[index].Username == creds[0] && users[index].Password == creds[1] {
			var token Token
			token.generate(users[index].Username)
			tokens[token.Username] = token.Secret
			json.NewEncoder(w).Encode(token.Secret)
			break
		}
	}

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

	secret := make([]byte, 8)
	rand.Read(secret)
	fmt.Println(secret)

	http.ListenAndServe(":8080", nil)
}
