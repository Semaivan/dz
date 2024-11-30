package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

// функция получения чисел от 1 до 6
func kubik(w http.ResponseWriter, r *http.Request) {
	a := rand.Intn(6) + 1
	response := []byte(strconv.Itoa(a))
	w.Write(response)
	return
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/play", kubik) //обработчик
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe() // запуск http сервера
}
