package main

import (
	"database/sql"
	"log"
)

type InvoiceDao interface {
	GetInvoices() ([]*Invoice, error)
	GetInvoiceById(id string) (*Invoice, error)
	CreateInvoice(invoice *Invoice) error
	UpdateInvoice(invoice *Invoice) error
	DeleteInvoice(id string) error
}

type InvoiceDaoImpl struct {
	conn *sql.DB
}

func NewInvoiceDao(conn *sql.DB) *InvoiceDaoImpl {
	return &InvoiceDaoImpl{conn: conn}
}

func (dao *InvoiceDaoImpl) GetInvoices() ([]*Invoice, error) {
	log.Println("Get all invoices")
	rows, err := dao.conn.Query("SELECT id, customer_id, purchaseDate FROM invoice")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invoices := []*Invoice{}
	for rows.Next() {
		invoice := new(Invoice)
		err := rows.Scan(&invoice.ID, &invoice.CustomerID, &invoice.PurchaseDate)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return invoices, nil

}

func (dao *InvoiceDaoImpl) GetInvoiceById(id string) (*Invoice, error) {
	log.Println("Get invoice by id")
	row := dao.conn.QueryRow("SELECT id, customer_id, purchaseDate FROM invoice WHERE id = $1", id)
	invoice := new(Invoice)
	err := row.Scan(&invoice.ID, &invoice.CustomerID, &invoice.PurchaseDate)
	if err != nil {
		return nil, err
	}
	return invoice, nil
}

func (dao *InvoiceDaoImpl) CreateInvoice(invoice *Invoice) error {
	log.Println("Create invoice")
	_, err := dao.conn.Exec("INSERT INTO invoice (customer_id, purchaseDate) VALUES ($1, $2)", invoice.CustomerID, invoice.PurchaseDate)
	if err != nil {
		return err
	}
	return nil
}

func (dao *InvoiceDaoImpl) UpdateInvoice(invoice *Invoice) error {
	log.Println("Update invoice")
	_, err := dao.conn.Exec("UPDATE invoice SET customer_id = $1, purchaseDate = $2 WHERE id = $3", invoice.CustomerID, invoice.PurchaseDate, invoice.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *InvoiceDaoImpl) DeleteInvoice(id string) error {
	log.Println("Delete invoice")
	_, err := dao.conn.Exec("DELETE FROM invoice WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
