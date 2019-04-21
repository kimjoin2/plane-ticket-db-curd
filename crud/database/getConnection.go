package database

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", getSourcePath())

	if err != nil {
		log.Println(err)
		return nil, errors.New(getConnectionErrorMsg())
	}

	return db, nil
}
