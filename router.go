package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

// O: A Router struct
// D: Will return a new Router struct
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// I: A http writer and a request
// D: This will find a Handler in Router struct by the url path and serve it
func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, methodExists, exists := r.FindHandler(req.URL.Path, req.Method)

	if !exists {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	if !methodExists {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	handler(w, req)
}

// I: A URL Path and a HTTP Method as a string
// O: A handler, and two booleans
// D: Will return the handler func if exists, else will return two booleans to
//    To indicate if the path or method exists
func (r Router) FindHandler(path, method string) (http.HandlerFunc, bool, bool) {
	_, exists := r.rules[path]
	handler, methodExists := r.rules[path][method]

	return handler, methodExists, exists
}
