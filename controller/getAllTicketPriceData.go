package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"planeTicketCrudService/crud"
)

func GetAllFlightDataController(w http.ResponseWriter, r *http.Request) {
	_ = r
	data, err := crud.GetAllFlightDataDAO()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		jsonString, _ := json.Marshal(data)
		fmt.Fprintln(w, string(jsonString))
	}
}
