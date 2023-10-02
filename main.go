package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	r := mux.NewRouter()

	conn, err := sql.Open("pgx", "postgres://acme:AcM3!@localhost:5432/acme?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	NewCustomerHandler(NewCustomerDao(conn), r)
	NewProductHandler(NewProductDao(conn), r)
	NewInvoiceHandler(NewInvoiceDao(conn), r)
	NewInvoiceItemHandler(NewInvoiceItemDao(conn), r)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
