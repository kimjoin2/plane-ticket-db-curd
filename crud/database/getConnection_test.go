package database

import (
	"plane-ticket-db-curd/crud"
	"testing"
)

func TestGetConnection(t *testing.T) {
	db, err := crud.GetConnection()
	if db == nil || err != nil {
		t.Error(err)
	}
}