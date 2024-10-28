package main

import (
	"fmt"
	"golang_api/handlers"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	server := constructServer()
	fmt.Println("Сервер localhost:8090 запущен и прослушивает входящие запросы...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("ошибка при запуске сервера.\n%v)", err)
	}
}

// constructServer возвращает инициализированный сервер типа *http.Server
func constructServer() *http.Server {
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