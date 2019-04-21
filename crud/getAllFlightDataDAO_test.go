package crud

import (
	"testing"
)

func TestGetAllFlightDataDAO(t *testing.T) {
	data, err := GetAllFlightDataDAO()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(data)
	}
}
