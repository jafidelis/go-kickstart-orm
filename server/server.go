package server

import (
	"net/http"

	"github.com/go-kickstart-orm/auth"

	"github.com/gorilla/mux"
)

//NewServer with dinamic routes
func NewServer() *mux.Router {
	mux := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		if route.RequireAuth {
			handler = auth.OAuthFilter(handler)
		}
		mux.
			Methods(route.Method).
			Name(route.Name).
			Path("/api" + route.Path).
			Handler(handler)
	}
	return mux
}
