package crud

import (
	"errors"
	"log"
	"planeTicketCrudService/crud/database"
	"planeTicketCrudService/dto"
	"planeTicketCrudService/utils"
	"time"
)

func GetAllFlightDataDAO() ([]dto.TicketPriceDTO, error) {

	db, err := database.GetConnection()
	if err != nil {
		log.Println(err)
		return nil, errors.New(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(getAllTicketDataSQL)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	var res []dto.TicketPriceDTO


	index := 0
	for rows.Next(){
		var data dto.TicketPriceDTO
		var count int32

		var baseDepartTime, targetArrivalTime, targetDepartTime, baseArrivalTime time.Time

		err := rows.Scan(
			&data.BaseAirport,
			&data.BaseNation,
			&data.TargetAirport,
			&data.TargetNation,
			&data.BaseAirline,
			&data.TargetAirline,
			&baseDepartTime,
			&targetArrivalTime,
			&targetDepartTime,
			&baseArrivalTime,
			&data.Price,
			&count)
		if err != nil {
			log.Println(err.Error())
			break
		}
		data.BaseDepartTime = utils.TimeToStringUntilMin(baseDepartTime)
		data.BaseArrivalTime = utils.TimeToStringUntilMin(baseArrivalTime)
		data.TargetDepartTime = utils.TimeToStringUntilMin(targetDepartTime)
		data.TargetArrivalTime = utils.TimeToStringUntilMin(targetArrivalTime)
		if res == nil {
			res = make([]dto.TicketPriceDTO, count)
		}
		res[index] = data
		index++
	}

	return res, nil
}
