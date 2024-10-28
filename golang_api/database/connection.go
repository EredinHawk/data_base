package database

import (
	"database/sql"
	"golang_api/lib"
)

/*
Connection возвращает экземпляр объекта sql.DB
с помощью которого будут осуществляться sql запросы.
*/
func Connection() (*sql.DB, error) {
	db, err := sql.Open("postgres", lib.ConnectionStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
