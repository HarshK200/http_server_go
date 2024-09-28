package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// NOTE: in memory storage i.e. imitating a database
var Users = []User{}

type Api struct {
	addr string
}

func (api *Api) getUserHanlder(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// encoding user struct to json
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(Users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error occured during json encoding users arr")
	}
}

func (api *Api) createUserHanlder(w http.ResponseWriter, req *http.Request) {
	var payload User

	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validateData(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := insertUser(&payload, &Users); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user registered successfuly"))
}

func validateData(payload *User) error {
	if payload.Name == "" {
		return errors.New("username is required")
	}
	if payload.Email == "" {
		return errors.New("email is required")
	}

	return nil
}

func insertUser(payload *User, users *[]User) error {

	for _, user := range *users {
		if user.Email == payload.Email {
			return errors.New("user with same email already exists")
		}
	}

	newUser := User{
		Name:  payload.Name,
		Email: payload.Email,
	}

	*users = append(*users, newUser)
	return nil
}
