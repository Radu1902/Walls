package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var users []User = []User{{"ionut", "qwer", "salut"}, {"mircea", "asdf", "buna"}}
var tokens map[string]string = make(map[string]string)

func (database *Database) init() error {
	userlist, err := database.db.Query("SELECT * FROM users")
	if err != nil {
		return err
	}
	fmt.Println(userlist)
	return err
}

// to do: 69 the map

func (database *Database) registerHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var credentials map[string]string
	decoder.Decode(&credentials)

	for _, user := range users {
		if credentials["username"] == user.Username {
			fmt.Println("Username already used")
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

	newUser := User{credentials["username"], credentials["password"], "Hello! this is my wall."}
	users = append(users, newUser)
	database.db.Query("INSERT INTO users (username, password, wall) VALUES (?, ?, ?)", newUser.Username, newUser.Password, newUser.Wall)

	// to do: db sync
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var credentials map[string]string
	decoder.Decode(&credentials)

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

func updateWallHandler(w http.ResponseWriter, r *http.Request) {
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

func getWall(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/user/"):]
	if username == "" {
		http.NotFound(w, r)
		return
	}
	for index := range users {
		if username == users[index].Username {
			json.NewEncoder(w).Encode(users[index].Wall)
			return
		}
	}
	http.NotFound(w, r)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "login.html")
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "edit.html")
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "search.html")
}
