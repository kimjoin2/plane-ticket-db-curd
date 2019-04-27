package crud

import "testing"

func TestInputFlightDataDAO(t *testing.T) {

	data := make(map[string]string)
	data[baseAirport] = "test_airport"
	data[targetAirport] = "test_airport2"
	data[baseDepart] = "1991/3/20 11:20"
	data[targetArrival] = "1992/3/20 13:13"
	data[baseAirline] = "test_airline"
	data[targetAirline] = "test_airline2"
	data[price] = "2000"

	if err := InputFlightDataDAO(data); err != nil {
		t.Fatal(err)
	} else {
		t.Log("success")
	}
}

