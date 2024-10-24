package db

import (
	"database/sql"
	"fmt"
	"golang_api/lib"

	_ "github.com/lib/pq"
)

// SearchFlights возвращает результат sql запроса к БД
//
// Поиск авиабилетов по параметрам
func SearchFlights(db *sql.DB, srh lib.Search) {
	reqStr := fmt.Sprintf("SELECT * FROM fligh")
	db.Query(reqStr,)
}