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

	result, err := database.GetFlightsRequest(params)
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
	if string(input) == "null"{
		http.Error(w, "Авиарейсов с такими параметрами отстуствуют", 404)
	}
	fmt.Fprintf(w, string(input))
}
