package controller

import (
	"net/http"
)

func InputFlightDataController(w http.ResponseWriter, r *http.Request) {
	_ = r
	//params := r.URL.Query()
	//for k, v := range params {
	//
	//}
	w.WriteHeader(http.StatusNoContent)
}
