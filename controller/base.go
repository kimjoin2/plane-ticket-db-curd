package controller

import (
	"fmt"
	"log"
	"net/http"
)

func BaseController(w http.ResponseWriter, r *http.Request) {
	_ = r
	_, err := fmt.Fprintln(w, "welcome!")
	if err != nil {
		log.Println(err)
	}
}
