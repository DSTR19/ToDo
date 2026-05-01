package core_http_server

import (
	"net/http"
)

type Route struct {
	Method  string       // GET, POST, PUT, DELETE, etc.
	Path    string       // e.g., "/users/{id}"
	Handler http.Handler // The handler function to execute when the route is matched
}

func NewRoute(
	method string,
	path string,
	handler http.Handler,
) Route {
	return Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}
