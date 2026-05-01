package users_transport_http

import (
	"net/http"

	core_http_server "github.com/DSTR19/ToDo/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	UsersService UsersService
}

type UsersService interface {
}

func NewUsersHTTPHandler(usersService UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		UsersService: usersService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: http.HandlerFunc(h.CreateUser),
		},
	}
}
