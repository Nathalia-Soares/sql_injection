package main

import (
    "log"
    "net/http"

    "controller"
)

func main() {
    controller.SetupRotas()
    log.Println("Servidor rodando na porta 8080")
    http.ListenAndServe(":8080", nil)
}