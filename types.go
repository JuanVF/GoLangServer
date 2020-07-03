package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// D: This will return a User Struct as a byte slice, will allow
//    the server the send a res as a JSON
func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
