package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/markusremplbauer/go-rest-api/api/data"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var p data.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.CreateProduct(data.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}
