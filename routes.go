package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	name       string
	method     string
	pattern    string
	handlefunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Name(route.name).
			Methods(route.method).
			Path(route.pattern).
			Handler(route.handlefunc)
	}

	return router
}

var routes = Routes{
	Route{"index",
		"GET",
		"/",
		Index},
	Route{"MovieList",
		"GET",
		"/peliculas",
		MovieList},
	Route{"MovieShow",
		"GET",
		"/pelicula/{id}",
		MovieShow},
}
