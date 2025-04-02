package produtos_vulneravel

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "host=localhost port=5432 user=your_user password=your_password dbname=your_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func getProdutosVulneravel(w http.ResponseWriter, r *http.Request) {
	categoria := r.URL.Query().Get("categoria")
	query := "SELECT * FROM products WHERE categoria = '" + categoria + "'"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var produtos []map[string]interface{}
	for rows.Next() {
		var id int
		var nome, categoria string
		if err := rows.Scan(&id, &name, &category); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		produtos = append(produtos, map[string]interface{}{"id": id, "nome": nome, "categoria": categoria})
	}

	json.NewEncoder(w).Encode(produtos)
}
