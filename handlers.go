package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

var users []User = []User{{"ionut", "qwer", "salut"}, {"mircea", "asdf", "buna"}}
var tokens map[string]string = make(map[string]string)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var credentials map[string]string
	decoder.Decode(&credentials)
	newUser := User{credentials["username"], credentials["password"], ""}
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
		fmt.Println(token.Secret)
		fmt.Println(tokens)
		fmt.Println(tokens[token.Secret])
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
	fmt.Println("sesh: ", sessionUser)
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

func loginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "edit.html")
}
