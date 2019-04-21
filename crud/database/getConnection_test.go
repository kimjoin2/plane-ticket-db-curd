package database

import (
	"planeTicketCrudService/crud"
	"testing"
)

func TestGetConnection(t *testing.T) {
	db, err := crud.GetConnection()
	if db == nil || err != nil {
		t.Error(err)
	}
}