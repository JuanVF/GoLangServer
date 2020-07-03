package main

//Here will be defined URL Paths and Middlewares to use by each path
func main() {
	server := NewServer(":3000")
	server.Handle("/", "GET", HandleRoot)
	server.Handle("/user", "POST", server.AddMiddleware(UserPostRequest, CheckAuth(), Logging()))
	server.Handle("/hi", "POST", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
