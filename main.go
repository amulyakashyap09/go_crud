package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Age      int32    `json:"age"`
	Contact  *Contact `json:"contact"`
}

type Contact struct {
	Mobile  string `json:"mobile"`
	Address string `json:"address"`
}

var users []User

func v1GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func v1GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func v1CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(100000))
	users = append(users, user)
	_ = json.NewEncoder(w).Encode(user)
}

func v1UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...) //sliced now book holds the particular element
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = params["id"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
}

func v1DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	//initialize router
	r := mux.NewRouter()

	//dummy data
	users = append(users, User{ID: "1", Username: "amulya1", Age: 26, Email: "amulya@gmail.com", Contact: &Contact{Mobile: "9559974779", Address: "Malad West, Mumbai, Maharashtra"}})
	users = append(users, User{ID: "2", Username: "amulya2", Age: 27, Email: "amulya2@gmail.com", Contact: &Contact{Mobile: "9559974779", Address: "Malad West, Mumbai, Maharashtra"}})
	users = append(users, User{ID: "3", Username: "amulya3", Age: 28, Email: "amulya3@gmail.com", Contact: &Contact{Mobile: "9559974779", Address: "Malad West, Mumbai, Maharashtra"}})

	r.HandleFunc("/v1/users", v1GetUsers).Methods("GET")
	r.HandleFunc("/v1/users/{id}", v1GetUser).Methods("GET")
	r.HandleFunc("/v1/users", v1CreateUser).Methods("POST")
	r.HandleFunc("/v1/users/{id}", v1UpdateUser).Methods("PUT")
	r.HandleFunc("/v1/users/{id}", v1DeleteUser).Methods("DELETE")

	fmt.Println("server running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r)))
}
