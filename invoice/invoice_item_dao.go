package invoice

import (
	"database/sql"
	"log"

	"github.com/rickCrz7/acme/utils"
)

type InvoiceItemDao interface {
	GetInvoiceItems() ([]*utils.InvoiceItem, error)
	GetInvoiceItemById(id string) (*utils.InvoiceItem, error)
	CreateInvoiceItem(invoiceItem *utils.InvoiceItem) error
	UpdateInvoiceItem(invoiceItem *utils.InvoiceItem) error
	DeleteInvoiceItem(id string) error
}

type InvoiceItemDaoImpl struct {
	conn *sql.DB
}

func NewInvoiceItemDao(conn *sql.DB) *InvoiceItemDaoImpl {
	return &InvoiceItemDaoImpl{conn: conn}
}

func (dao *InvoiceItemDaoImpl) GetInvoiceItems() ([]*utils.InvoiceItem, error) {
	log.Println("Get all invoiceItems")
	rows, err := dao.conn.Query("SELECT id, invoice_id, product_id, quantity, price FROM invoice_item order by id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	invoiceItems := []*utils.InvoiceItem{}
	for rows.Next() {
		invoiceItem := new(utils.InvoiceItem)
		err := rows.Scan(&invoiceItem.ID, &invoiceItem.InvoiceID, &invoiceItem.ProductID, &invoiceItem.Quantity, &invoiceItem.Price)
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

func (dao *InvoiceItemDaoImpl) GetInvoiceItemById(id string) (*utils.InvoiceItem, error) {
	log.Println("Get invoiceItem by id")
	row := dao.conn.QueryRow("SELECT id, invoice_id, product_id, quantity, price FROM invoice_item WHERE id = $1", id)
	invoiceItem := new(utils.InvoiceItem)
	err := row.Scan(&invoiceItem.ID, &invoiceItem.InvoiceID, &invoiceItem.ProductID, &invoiceItem.Quantity, &invoiceItem.Price)
	if err != nil {
		return nil, err
	}
	return invoiceItem, nil
}

func (dao *InvoiceItemDaoImpl) CreateInvoiceItem(invoiceItem *utils.InvoiceItem) error {
	log.Println("Create invoiceItem")
	_, err := dao.conn.Exec("INSERT INTO invoice_item (invoice_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)", invoiceItem.InvoiceID, invoiceItem.ProductID, invoiceItem.Quantity, invoiceItem.Price)
	if err != nil {
		return err
	}
	return nil
}

func (dao *InvoiceItemDaoImpl) UpdateInvoiceItem(invoiceItem *utils.InvoiceItem) error {
	log.Println("Update invoiceItem")
	_, err := dao.conn.Exec("UPDATE invoice_item SET invoice_id = $1, product_id = $2, quantity = $3, price = $4 WHERE id = $5", invoiceItem.InvoiceID, invoiceItem.ProductID, invoiceItem.Quantity, invoiceItem.Price , invoiceItem.ID)
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
