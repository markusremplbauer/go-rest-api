package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/markusremplbauer/go-rest-api/api/data"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	p := data.Product{ID: id}
	if err := p.DeleteProduct(data.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
