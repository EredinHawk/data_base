package database

import (
	"database/sql"
	"golang_api/lib"
)

var flights []lib.Flights

func SendFlightsRequest(s *lib.Search) ([]lib.Flights, error) {
	db, err := Connection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if s.City_depart == "null" && s.City_dest == "null" && s.Date == "null" {
		r, err := db.Query(SelectAllFlights)
		if err != nil {
			return nil, err
		}
		if err := scanStruct(r); err != nil {
			return nil, err
		}
	} else {
		r, err := db.Query(SelectParamFlights, s.City_depart, s.City_dest, s.Date)
		if err != nil {
			return nil, err
		}
		err = scanStruct(r)
		if err != nil {
			return nil, err
		}
	}
	return flights, nil
}

func scanStruct(r *sql.Rows) error {
	for r.Next() {
		var f lib.Flights
		err := r.Scan(&f.F_number, &f.F_time, &f.Al_name, &f.Ap_Name_depart,
			&f.Ap_Name_dest, &f.City_depart, &f.City_dest, &f.F_cost)
		if err != nil {
			return err
		}
		flights = append(flights, f)
	}
	return nil
}
