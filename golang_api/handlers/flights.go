package handlers

import (
	"fmt"
	aviadb "golang_api/db"
	"golang_api/encoding"
	"net/http"
)

// GetFlights возвращает список авиарейсов в формате JSON.
func GetFlights(w http.ResponseWriter, r *http.Request) {
	body, err := encoding.ScanBody(r)
	if err != nil {
		fmt.Fprintf(w, "GetFlights: ошибка при сканировании тела http запроса\n%v", err)
	}

	db, err := aviadb.Connection()
	if err != nil {
		fmt.Fprintf(w, "GetFlights: ошибка при подключении к БД\n%v", err)
	}

	aviadb.SearchFlights(db, body)
}
