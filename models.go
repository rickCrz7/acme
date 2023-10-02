package main

import "time"

type Customer struct {
	ID   string
	Name string
}

type Product struct {
	ID    string
	Name  string
	Price string
}

type Invoice struct {
	ID       string
	CustomerID string
	PurchaseDate time.Time
}

type InvoiceItem struct {
	ID       string
	InvoiceID string
	ProductID string
	Quantity  int
}