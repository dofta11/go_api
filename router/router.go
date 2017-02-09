package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"fmt"
	"go_api/logger"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		fmt.Print(" 1 => ")
		fmt.Print(routes)
		fmt.Println()
		fmt.Print(" 2 => ")
		fmt.Print(route)
		fmt.Println()

		var handler http.Handler

		handler = route.HandlerFunc
		fmt.Print(" 3 => ")
		fmt.Print(handler)
		fmt.Println()
		handler = logger.Logger(handler, route.Name)
		fmt.Print(" 4 => ")
		fmt.Print(handler)
		fmt.Println()

		fmt.Print(" Method => ")
		fmt.Print(route.Method)
		fmt.Println()
		fmt.Print(" Pattern => ")
		fmt.Print(route.Pattern)
		fmt.Println()
		fmt.Print(" Name => ")
		fmt.Print(route.Name)
		fmt.Println()

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
