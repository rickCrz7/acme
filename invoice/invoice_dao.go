package invoice

import (
	"database/sql"
	"log"

	"github.com/rickCrz7/acme/utils"
)

type InvoiceDao interface {
	GetInvoices() ([]*utils.Invoice, error)
	GetInvoiceById(id string) (*utils.Invoice, error)
	CreateInvoice(invoice *utils.Invoice) error
	UpdateInvoice(invoice *utils.Invoice) error
	DeleteInvoice(id string) error
}

type InvoiceDaoImpl struct {
	conn *sql.DB
}

func NewInvoiceDao(conn *sql.DB) *InvoiceDaoImpl {
	return &InvoiceDaoImpl{conn: conn}
}

func (dao *InvoiceDaoImpl) GetInvoices() ([]*utils.Invoice, error) {
	log.Println("Get all invoices")
	rows, err := dao.conn.Query("SELECT id, customer_id, purchaseDate FROM invoice ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invoices := []*utils.Invoice{}
	for rows.Next() {
		invoice := new(utils.Invoice)
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

func (dao *InvoiceDaoImpl) GetInvoiceById(id string) (*utils.Invoice, error) {
	log.Println("Get invoice by id")
	row := dao.conn.QueryRow("SELECT id, customer_id, purchaseDate FROM invoice WHERE id = $1", id)
	invoice := new(utils.Invoice)
	err := row.Scan(&invoice.ID, &invoice.CustomerID, &invoice.PurchaseDate)
	if err != nil {
		return nil, err
	}
	return invoice, nil
}

func (dao *InvoiceDaoImpl) CreateInvoice(invoice *utils.Invoice) error {
	log.Println("Create invoice")
	_, err := dao.conn.Exec("INSERT INTO invoice (customer_id, purchaseDate) VALUES ($1, $2)", invoice.CustomerID, invoice.PurchaseDate)
	if err != nil {
		return err
	}
	return nil
}

func (dao *InvoiceDaoImpl) UpdateInvoice(invoice *utils.Invoice) error {
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
