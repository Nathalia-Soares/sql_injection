package produtos_protegido

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {

	var err error

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	log.Printf("Conectando ao banco em %s:%s como %s no banco %s", host, port, user, dbname)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
}

func GetProdutosProtegido(w http.ResponseWriter, r *http.Request) {
	categoria := r.URL.Query().Get("categoria")
	query := "SELECT * FROM produtos WHERE categoria = $1"
	rows, err := db.Query(query, categoria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var produtos []map[string]interface{}
	for rows.Next() {
		var id int
		var nome, categoria, img_url string
		if err := rows.Scan(&id, &nome, &categoria, &img_url); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		produtos = append(produtos, map[string]interface{}{"id": id, "nome": nome, "categoria": categoria, "img_url": img_url})
	}

	json.NewEncoder(w).Encode(produtos)
}
