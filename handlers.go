package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

var users []User = []User{{"ionut", "qwer", []byte("salut")}, {"mircea", "asdf", []byte("buna")}}
var tokens map[string]string = make(map[string]string)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body string
	decoder.Decode(&body)
	creds := strings.Fields(body)
	newUser := User{creds[0], creds[1], []byte("")}
	users = append(users, newUser)
	// to do: db sync
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var credentials map[string]string
	decoder.Decode(&credentials)
	fmt.Println(reflect.TypeOf(credentials))
	fmt.Println(credentials["username"], credentials["password"])

	encoder := json.NewEncoder(w)
	var authSuccess bool = false
	var token Token

	for index := range users {
		if users[index].Username == credentials["username"] && users[index].Password == credentials["password"] {
			token.generate(users[index].Username)
			tokens[token.Secret] = token.Username
			authSuccess = true
			break
		}
	}
	if authSuccess {
		encoder.Encode(token.Secret)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}

func getWall(w http.ResponseWriter, r *http.Request) {
	var secret string
	json.NewDecoder(r.Body).Decode(&secret)
	fmt.Println(secret)
	var sessionUser string = tokens[secret]
	if sessionUser != "" {
		for index := range users {
			if sessionUser == users[index].Username {
				json.NewEncoder(w).Encode(users[index].Wall)
			}
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}

func getUser(w http.ResponseWriter, r *http.Request) {
	usr := users[1]
	fmt.Println(usr)
	json.NewEncoder(w).Encode(usr)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "edit.html")
}
