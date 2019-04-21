package main

import (
	"log"
	"net/http"
	"planeTicketCrudService/controller"
	"planeTicketCrudService/webServerConfig"
)

func main() {
	http.HandleFunc("/", controller.BaseController)
	http.HandleFunc("/getAll", controller.GetAllFlightDataController)
	log.Println("server started!")
	err := http.ListenAndServe(webServerConfig.GetServerPortData(), nil)
	if err != nil {
		log.Println(err)
	}
}
