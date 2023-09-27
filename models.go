package main

import "time"

type Customer struct {
	ID   int
	Name string
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

type Invoice struct {
	ID       int
	CustomerID int
	PurchaseDate time.Time
}

type InvoiceItem struct {
	ID       int
	InvoiceID int
	ProductID int
	Quantity  int
}