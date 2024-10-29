package database

import (
	"database/sql"
	"fmt"
	"golang_api/lib"
	"strings"
)

// GetFlightsRequest осуществляет SQL запрос к БД, который возвращает 
// список авиарейсов. Если параметры не важны, то вернутся все авиарейсы.
// Если есть входящие параметры, то вернутся конкретные авиарейсы, подходящие
// условию.
func GetFlightsRequest(s *lib.Search) ([]lib.Flights, error) {
	db, err := Connection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var flights []lib.Flights

	if s.City_depart == "null" && s.City_dest == "null" && s.Date == "null" {
		rows, err := db.Query(SelectAllFlights)
		if err != nil {
			return nil, err
		}

		flights, err = scanStruct(rows)
		if err != nil {
			return nil, err
		}

	} else {
		date, err := splitDate(s.Date)
		if err != nil {
			return nil, err
		}

		rows, err := db.Query(SelectParamFlights, s.City_depart, s.City_dest, date[0], date[1], date[2])
		if err != nil {
			return nil, err
		}

		flights, err = scanStruct(rows)
		if err != nil {
			return nil, err
		}
	}

	return flights, nil
}

// scanStruct копирует строки результирующего набора в поля экземпляра Flights
func scanStruct(r *sql.Rows) ([]lib.Flights, error) {
	var flights []lib.Flights
	for r.Next() {
		var f lib.Flights
		err := r.Scan(&f.F_number, &f.F_time, &f.Al_name, &f.Ap_Name_depart,
			&f.Ap_Name_dest, &f.City_depart, &f.City_dest, &f.F_cost)
		if err != nil {
			return nil, err
		}
		flights = append(flights, f)
	}
	return flights, nil
}

// splitDate разделяет строку формата "2024-02-24" на срез []{2024, 02, 24}
func splitDate(s string) ([]string, error) {
	result := strings.Split(s, "-")
	if len(result) != 3 {
		return nil, fmt.Errorf("недопустимое значение даты, введите дату в формате 1212-12-12")
	}
	return result, nil
}
