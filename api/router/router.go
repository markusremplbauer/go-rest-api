package router

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		handler := Logger(route.HandlerFunc)

		router.Methods(route.Method).Path(route.Pattern).Handler(handler)
	}

	return router
}
