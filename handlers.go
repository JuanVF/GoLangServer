package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// I: A http writer and a request
// D: The / page result
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the main page :)")
}

// I: A http writer and a request
// D: The /hi page result
func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is another end point :D")
}

// I: A http writer and a request
// D: The /user page result, this will return the User struct received by request
//    As a JSON
func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user User

	err := decoder.Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)

		return
	}
	res, err := user.ToJson()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(res)
}
