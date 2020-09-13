package main

import (
	"log"
	"net/http"
	"os"

	"github.com/markusremplbauer/go-rest-api/api/data"
	"github.com/markusremplbauer/go-rest-api/api/router"
)

func main() {
	router := router.NewRouter()

	data.Initialize(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	log.Fatal(http.ListenAndServe(":8080", router))
}
