package crud

import (
	"log"
	"plane-ticket-db-curd/crud/database"
	"strconv"
	"strings"
	"time"
)

func InputFlightDataDAO(data map[string]string) error {
	db, err := database.GetConnection()
	if err != nil {
		log.Println(err)
		return err
	}
	defer db.Close()

	// TODO : get airline index
	var baseAirlineId int
	baseAirlineRow := db.QueryRow(getAirlineListWithWhereName, data[baseAirline])
	baseAirlineRow.Scan(&baseAirlineId)

	var targetAirlineId int
	targetAirlineRow := db.QueryRow(getAirlineListWithWhereName, data[targetAirline])
	targetAirlineRow.Scan(&targetAirlineId)

	// TODO : get airport index
	var baseAirportId int
	baseAirportRow := db.QueryRow(getAirportListWithWhereCode, data[baseAirport])
	baseAirportRow.Scan(&baseAirportId)

	var targetAirportId int
	targetAirportRow := db.QueryRow(getAirportListWithWhereCode, data[targetAirport])
	targetAirportRow.Scan(&targetAirportId)

	// TODO : make date instance
	baseDepartTimeStringSplit := strings.Split(data[baseDepart], " ")
	baseDepartTimeDaySplit := strings.Split(baseDepartTimeStringSplit[0], "/")
	baseDepartTimeTimeSplit := strings.Split(baseDepartTimeStringSplit[1], ":")

	baseDepartYear, err := strconv.Atoi(baseDepartTimeDaySplit[0])
	if err != nil {
		return err
	}
	baseDepartMonth, err := strconv.Atoi(baseDepartTimeDaySplit[1])
	if err != nil {
		return err
	}
	baseDepartDay, err := strconv.Atoi(baseDepartTimeDaySplit[2])
	if err != nil {
		return err
	}

	baseDepartHour, err := strconv.Atoi(baseDepartTimeTimeSplit[0])
	if err != nil {
		return err
	}
	baseDepartMinutes, err := strconv.Atoi(baseDepartTimeTimeSplit[1])
	if err != nil {
		return err
	}

	baseDepartTime := time.Date(
		baseDepartYear,
		time.Month(baseDepartMonth),
		baseDepartDay,
		baseDepartHour,
		baseDepartMinutes,
		0,
		0,
		time.Local)

	targetArrivalTimeStringSplit := strings.Split(data[targetArrival], " ")
	targetArrivalTimeDaySplit := strings.Split(targetArrivalTimeStringSplit[0], "/")
	targetArrivalTimeTimeSplit := strings.Split(targetArrivalTimeStringSplit[1], ":")

	targetArrivalYear, err := strconv.Atoi(targetArrivalTimeDaySplit[0])
	if err != nil {
		return err
	}
	targetArrivalMonth, err := strconv.Atoi(targetArrivalTimeDaySplit[1])
	if err != nil {
		return err
	}
	targetArrivalDay, err := strconv.Atoi(targetArrivalTimeDaySplit[2])
	if err != nil {
		return err
	}

	targetArrivalHour, err := strconv.Atoi(targetArrivalTimeTimeSplit[0])
	if err != nil {
		return err
	}
	arrivalMinutes := strings.Split(targetArrivalTimeTimeSplit[1], "+")
	if len(arrivalMinutes) == 2 {
		targetArrivalDay++
	}
	targetArrivalMinutes, err := strconv.Atoi(arrivalMinutes[0])
	if err != nil {
		return err
	}

	targetArrivalTime := time.Date(
		targetArrivalYear,
		time.Month(targetArrivalMonth),
		targetArrivalDay,
		targetArrivalHour,
		targetArrivalMinutes,
		0,
		0,
		time.Local)

	// TODO : insert data
	_, err = db.Exec(inputTicketData,
		baseAirportId, targetAirportId,
		baseAirlineId, baseAirlineId,
		baseDepartTime, targetArrivalTime,
		data[price])
	if err != nil {
		return err
	}

	return nil
}
