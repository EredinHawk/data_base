package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	AllFlights = `SELECT
	FLIGHT.F_NUMBER,
	FLIGHT.F_TIME,
	AIRLINE.AL_NAME,
	AP1.AP_NAME,
	AP2.AP_NAME,
	C1.C_NAME,
	C2.C_NAME,
	FLIGHT.F_COST
FROM
	FLIGHT
	JOIN AIRLINE 		ON F_AIRLINE_ID	= AL_ID
	JOIN ROUTE 			ON F_ROUTE 		= R_ID
	JOIN AIRPORT AP1 	ON AP1.AP_ID 	= ROUTE.R_AIRPORT_DEPART
	JOIN AIRPORT AP2 	ON AP2.AP_ID 	= ROUTE.R_AIRPORT_DEST
	JOIN CITY C1 		ON C1.C_ID 		= AP1.AP_CITY_ID
	JOIN CITY C2 		ON C2.C_ID 		= AP2.AP_CITY_ID;`
)

var (
	BodyScanError  string = "GetFlights: ошибка при сканировании тела http запроса\n%v"
	DBConnectError string = "GetFlights: ошибка при подключении к БД\n%v"
)

// Парметры поиска авиабилетов
type Search struct {
	City_depart string `json:"city_depart"`
	City_dest   string `json:"city_dest"`
	Time        string `json:"time"`
}

// Авиабилеты
type Flights struct {
	F_id         int    `json:"id"`
	F_number     string `json:"number"`
	F_airline_id int    `json:"airline_id"`
	F_time       string `json:"time"`
	F_cost       string `json:"cost"`
}

// GetFlights возвращает список авиарейсов в формате JSON.
func GetFlights(w http.ResponseWriter, r *http.Request) {
	db, err := Connection()
	if err != nil {
		fmt.Fprintf(w, DBConnectError, err)
	}
	defer db.Close()

	body, err := ScanBody(r)
	if err != nil {
		fmt.Fprintf(w, BodyScanError, err)
	}

}

/*
Connection возвращает экземпляр объекта sql.DB
с помощью которого будут осуществляться sql запросы
*/
func Connection() (*sql.DB, error) {
	connStr := "user=admin dbname=aviadb password=password port=5432 sslmode=prefer"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return &sql.DB{}, err
	}
	return db, nil
}

/*
ScanBody сканирует тело входящего http запроса и декодирует его строки во внутреннюю структуру Search,
которая содержит в себе параметры поиска авиабилетов.
*/
func ScanBody(r *http.Request) (Search, error) {
	search := Search{}
	err := json.NewDecoder(r.Body).Decode(&search)
	if err != nil {
		return search, err
	}
	return search, nil
}
