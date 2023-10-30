package utils

import "time"

type Customer struct {
	ID   string "json: id"
	Name string "json: name"
}

type Product struct {
	ID    string "json: id"
	Name  string "json: name"
	Price string "json: price"
}

type Invoice struct {
	ID           string    "json: id"
	CustomerID   string    "json: customer_id"
	PurchaseDate time.Time "json: purchase_date"
	Items        []Item    "json: items"
}

type InvoiceItem struct {
	ID        string "json: id"
	InvoiceID string "json: invoice_id"
	ProductID string "json: product_id"
	Quantity  int    "json: quantity"
}

type Item struct {
	ID       string  "json: id"
	Quantity int     "json: quantity"
	Product  Product "json: product"
}

type ProductReport struct {
	ID       string "json: id"
	Name     string "json: name"
	Price    string "json: price"
	Quantity int    "json: quantity"
}

type TotalSold struct {
	ID        string "json: id"
	Name      string "json: name"
	Price     string "json: price"
	TotalSold string "json: price"
}

type CustomerReport struct {
	ID         string "json: id"
	Name       string "json: name"
	Quantity   string "json: price"
	TotalSales string "json: total_sales"
}

type DateReport struct {
	Date       string "json: date"
	TotalSales string "json: total_sales"
}
