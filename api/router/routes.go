package router

import (
	"net/http"

	"github.com/markusremplbauer/go-rest-api/api/handlers"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"GET", "/products", handlers.GetProducts},
	Route{"GET", "/products/{id:[0-9]+}", handlers.GetProduct},
	Route{"POST", "/products", handlers.CreateProduct},
	Route{"PUT", "/products/{id:[0-9]+}", handlers.UpdateProduct},
	Route{"DELETE", "/products/{id:[0-9]+}", handlers.DeleteProduct},
}
