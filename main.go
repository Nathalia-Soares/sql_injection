package main

import (
	"log"
	"net/http"

	"github.com/Nathalia-Soares/sql_injection/controller"
	"github.com/Nathalia-Soares/sql_injection/middlewares/cors"
)

func main() {
	controller.SetupController()

	handler := cors.EnableCORS(http.DefaultServeMux)

	log.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", handler)
}
