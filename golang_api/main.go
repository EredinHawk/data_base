package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"golang_api/server"
)

func main() {
	server := server.ConstructServer()
	fmt.Println("Сервер localhost:8090 запущен и прослушивает входящие запросы...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("ошибка при запуске сервера.\n%v)", err)
	}
}
