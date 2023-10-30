package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rickCrz7/acme/customer"
	"github.com/rickCrz7/acme/invoice"
	"github.com/rickCrz7/acme/product"
	"github.com/rickCrz7/acme/reports"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	r := mux.NewRouter()

	conn, err := sql.Open("pgx", "postgres://postgres:0241!@localhost:5432/acme?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	customer.NewCustomerHandler(customer.NewCustomerDao(conn), r)
	product.NewProductHandler(product.NewProductDao(conn), r)
	invoice.NewInvoiceHandler(invoice.NewInvoiceDao(conn), r)
	invoice.NewInvoiceItemHandler(invoice.NewInvoiceItemDao(conn), r)
	reports.NewReportsHandler(reports.NewReportsDao(conn), r)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
