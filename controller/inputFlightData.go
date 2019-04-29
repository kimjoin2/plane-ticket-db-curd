package controller

import (
	"log"
	"net/http"
	"plane-ticket-db-curd/crud"
)

func InputFlightDataController(w http.ResponseWriter, r *http.Request) {
	_ = r
	getParams := r.URL.Query()
	mapParams := make(map[string]string)

	for k, v := range getParams {
		mapParams[k] = v[0]
	}

	err := crud.InputFlightDataDAO(mapParams)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
