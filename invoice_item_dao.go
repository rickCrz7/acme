package main

import (
	"database/sql"
	"log"
)

type InvoiceItemDao interface {
	GetInvoiceItems() ([]*InvoiceItem, error)
	GetInvoiceItemById(id string) (*InvoiceItem, error)
	CreateInvoiceItem(invoiceItem *InvoiceItem) error
	UpdateInvoiceItem(invoiceItem *InvoiceItem) error
	DeleteInvoiceItem(id string) error
}

type InvoiceItemDaoImpl struct {
	conn *sql.DB
}

func NewInvoiceItemDao(conn *sql.DB) *InvoiceItemDaoImpl {
	return &InvoiceItemDaoImpl{conn: conn}
}


func (dao *InvoiceItemDaoImpl) GetInvoiceItems() ([]*InvoiceItem, error) {
	log.Println("Get all invoiceItems")
	rows, err := dao.conn.Query("SELECT id, invoice_id, product_id, quantity FROM invoice_item")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invoiceItems := []*InvoiceItem{}
	for rows.Next() {
		invoiceItem := new(InvoiceItem)
		err := rows.Scan(&invoiceItem.ID, &invoiceItem.InvoiceID, &invoiceItem.ProductID, &invoiceItem.Quantity)
		if err != nil {
			return nil, err
		}
		invoiceItems = append(invoiceItems, invoiceItem)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return invoiceItems, nil
}

func (dao *InvoiceItemDaoImpl) GetInvoiceItemById(id string) (*InvoiceItem, error) {
	log.Println("Get invoiceItem by id")
	row := dao.conn.QueryRow("SELECT id, invoice_id, product_id, quantity FROM invoice_item WHERE id = $1", id)
	invoiceItem := new(InvoiceItem)
	err := row.Scan(&invoiceItem.ID, &invoiceItem.InvoiceID, &invoiceItem.ProductID, &invoiceItem.Quantity)
	if err != nil {
		return nil, err
	}
	return invoiceItem, nil
}

func (dao *InvoiceItemDaoImpl) CreateInvoiceItem(invoiceItem *InvoiceItem) error {
	log.Println("Create invoiceItem")
	_, err := dao.conn.Exec("INSERT INTO invoice_item (invoice_id, product_id, quantity) VALUES ($1, $2, $3)", invoiceItem.InvoiceID, invoiceItem.ProductID, invoiceItem.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func (dao *InvoiceItemDaoImpl) UpdateInvoiceItem(invoiceItem *InvoiceItem) error {
	log.Println("Update invoiceItem")
	_, err := dao.conn.Exec("UPDATE invoice_item SET invoice_id = $1, product_id = $2, quantity = $3 WHERE id = $4", invoiceItem.InvoiceID, invoiceItem.ProductID, invoiceItem.Quantity, invoiceItem.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *InvoiceItemDaoImpl) DeleteInvoiceItem(id string) error {
	log.Println("Delete invoiceItem")
	_, err := dao.conn.Exec("DELETE FROM invoice_item WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
