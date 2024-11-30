package main

import (
	configmail "MyDz/3-validation-api/configs"
	auth "MyDz/3-validation-api/internal/auth"
	"fmt"
	"net/http"
)

func main() {
	conf := configmail.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf})
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server started at http://localhost:8081")
	server.ListenAndServe() // запуск http сервера

}
