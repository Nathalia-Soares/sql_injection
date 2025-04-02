package controller

import (
    "net/http"
    "services/produtos_protegido"
    "services/produtos_vulneravel"
)

func SetupController() {
    http.HandleFunc("/produtos_vulneravel", produtos_vulneravel.GetProdutosVulneravel)
    http.HandleFunc("/produtos_protegido", produtos_protegido.GetProdutosProtegido)
}