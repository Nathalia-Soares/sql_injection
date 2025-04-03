package main

import (
	"log"
	"net/http"

	"github.com/Nathalia-Soares/sql_injection/controller"
)

func main() {
	controller.SetupController()
	log.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8081", nil)
}
