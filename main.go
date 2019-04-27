package main

import (
	"log"
	"net/http"
	"planeTicketCrudService/controller"
	"planeTicketCrudService/webServerConfig"
	"time"
)

func main() {
	//jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	const location = "Asia/Tokyo"
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc

	http.HandleFunc("/", controller.BaseController)
	http.HandleFunc("/getAll", controller.GetAllFlightDataController)
	http.HandleFunc("/input", controller.InputFlightDataController)
	log.Println("server started!")
	err = http.ListenAndServe(webServerConfig.GetServerPortData(), nil)
	if err != nil {
		log.Println(err)
	}
}
