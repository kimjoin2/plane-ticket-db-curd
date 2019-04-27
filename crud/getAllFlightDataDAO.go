package crud

import (
	"errors"
	"log"
	"plane-ticket-db-curd/crud/database"
	"plane-ticket-db-curd/dto"
	"plane-ticket-db-curd/utils"
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

	res := make([]dto.TicketPriceDTO, 0)

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
			&data.Price,
			&count,
			&targetDepartTime,
			&baseArrivalTime)
		_ = err
		//if err != nil {
		//	log.Println(err.Error(), "error!!!!")
		//	break
		//}
		data.BaseDepartTime = utils.TimeToStringUntilMin(baseDepartTime)
		data.BaseArrivalTime = utils.TimeToStringUntilMin(baseArrivalTime)
		data.TargetDepartTime = utils.TimeToStringUntilMin(targetDepartTime)
		data.TargetArrivalTime = utils.TimeToStringUntilMin(targetArrivalTime)
		res = append(res, data)
	}

	return res, nil
}
