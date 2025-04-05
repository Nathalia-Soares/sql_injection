package controller

import (
	"net/http"

	"github.com/Nathalia-Soares/sql_injection/services/produtos_protegido"
	"github.com/Nathalia-Soares/sql_injection/services/produtos_vulneravel"
)

func SetupController() {
	http.HandleFunc("/produtos_vulneravel", produtos_vulneravel.GetProdutosVulneravel)
	http.HandleFunc("/produtos_protegido", produtos_protegido.GetProdutosProtegido)
}
