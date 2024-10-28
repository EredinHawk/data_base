package handlers

import (
	"fmt"
	"golang_api/database"
	"golang_api/encoding"
	"golang_api/lib"
	"net/http"
)

// GetFlights возвращает список авиарейсов в формате JSON.
func GetFlights(w http.ResponseWriter, r *http.Request) {
	params, err := encoding.ScanBody(r)
	if err != nil {
		http.Error(w, fmt.Sprintf(lib.ScanBodyError, err), 400)
		return
	}

	result, err := database.SendFlightsRequest(params)
	if err != nil {
		http.Error(w, fmt.Sprintf(lib.QueryBDError, err), 400)
		return
	}

	input, err := encoding.EncodingJSON(result)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf(lib.MarshalError, err), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(input))
}
