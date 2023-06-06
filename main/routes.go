package main

import (
	mux "github.com/julienschmidt/httprouter"
)

type Route struct {
	Method string
	Path   string
	Handle mux.Handle
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/",
		Index,
	},
	Route{
		"GET",
		"/tracks",
		TrackIndex,
	},
	Route{
		"GET",
		"/tracks/:trackId",
		TrackShow,
	},
	Route{
		"POST",
		"/tracks",
		TrackCreate,
	},
}
