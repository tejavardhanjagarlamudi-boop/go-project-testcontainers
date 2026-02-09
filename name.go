package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	connSTR := "postgres://postgres:bobby478@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connSTR)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
		return
	}
	log.Println("db connected successfully")

	createTable(db)

	// start HTTP server
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello MAN what's up")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createTable(db *sql.DB) {
	query := `
CREATE TABLE IF NOT EXISTS product (
	id SERIAL PRIMARY KEY,
	name TEXT,
	price FLOAT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`

	

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Table created successfully")
}
