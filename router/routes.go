package router

import (
	"net/http"
	"go_api/handler"
	"go_api/controller"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"LoginCheck",
		"POST",
		"/member/login_check",
		controller.LoginCheck,
	},

	Route{
		"Index",
		"GET",
		"/",
		handler.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		handler.TodoIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		handler.TodoCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		handler.TodoShow,
	},
}