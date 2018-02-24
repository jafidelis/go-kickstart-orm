package server

import (
	"net/http"

	"github.com/go-kickstart-orm/api/endpoint"
)

type Route struct {
	RequireAuth bool
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

var routes = [...]Route{
	Route{true, "SaveUser", "POST", "/user/save", endpoint.SaveUser},
	Route{true, "GetAllUser", "GET", "/user/all", endpoint.GetAllUser},
	Route{false, "Login", "POST", "/login", endpoint.Login},
}
