package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fjgmelloni/fullcycle/multithreading/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cep/", handlers.CepHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
