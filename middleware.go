package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// O: A Middleware struct
// D: This is a mock func, just to test if a Middleware will work
func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			flag := true

			fmt.Println("Checking authentication")

			if flag {
				f(w, req)
			} else {
				return
			}
		}
	}
}

// O: A Middleware struct
// D: This is a mock func, just to test if a Middleware will work
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()

			defer func() {
				log.Println(req.URL.Path, time.Since(start))
			}()

			f(w, req)
		}
	}
}
