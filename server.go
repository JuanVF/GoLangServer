package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

// I: Server port as a string
// O: A Server struct
// D: Creates a new server struct
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

// O: An error
// D: This will start the server with the port specified in NewServer()
func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)

	if err != nil {
		return err
	}

	return nil
}

// I: The url path and HTTP Method as a string, and a handlerFunc
// D: This will add the handler to the Router struct, also create it in Router
//    if not exists
func (s *Server) Handle(path, method string, handler http.HandlerFunc) {
	_, exists := s.router.rules[path]

	if !exists {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}

	s.router.rules[path][method] = handler
}

// I: A HandlerFunc, and n Middlewares funcs
// O: A HandlerFunc
// D: This will add middlewares to a URL Path, also, will execute all of them before execute
//    the main HandlerFunc
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}
