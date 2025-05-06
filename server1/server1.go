package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// регистрируем маршрутизатор handler который будет обрабатывать все запросы к сайту
	http.HandleFunc("/", handler)

	//стартуем сервер
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// функция для обработки запросов
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
