package server

import (
	"golang_api/handlers"
	"net/http"
	"time"
)

// ConstructServer возвращает инициализированный сервер типа *http.Server
func ConstructServer() *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("GET /flights", handlers.GetFlights)

	srv := &http.Server{
		Addr:         "localhost:8090",
		Handler:      router, // HTTP мультиплексор, или по другому роутер
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return srv
}
