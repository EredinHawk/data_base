package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Connection возвращает экземпляр объекта sql.DB
// с помощью которого будут осуществляться sql запросы
func Connection() (*sql.DB, error) {
	connStr := "user=admin dbname=aviadb password=password port=5432 sslmode=prefer"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return &sql.DB{}, err
    }
    return db, nil
}