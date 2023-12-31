package reports

import (
	"database/sql"
	"log"

	"github.com/rickCrz7/acme/utils"
)

type ReportsDao interface {
	MostSoldProducts() ([]*utils.ProductReport, error)
	TotalSalesByProduct() ([]*utils.TotalSold, error)
	TotalSalesByCustomers() ([]*utils.CustomerReport, error)
	ReportsByStatus() ([]*utils.StatusReport, error)
}

type ReportsDaoImpl struct {
	conn *sql.DB
}

func NewReportsDao(conn *sql.DB) *ReportsDaoImpl {
	return &ReportsDaoImpl{conn: conn}
}

func (dao *ReportsDaoImpl) MostSoldProducts() ([]*utils.ProductReport, error) {
	log.Println("Get most sold products")
	rows, err := dao.conn.Query(`SELECT p.id, p.name, p.price, SUM(it.quantity) AS quantity 
	FROM invoice_item it INNER JOIN product p ON it.product_id = p.id 
	where it.invoice_id in (select id from invoice where status = true)
	GROUP BY p.id, p.name, p.price 
	ORDER BY quantity DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*utils.ProductReport{}
	for rows.Next() {
		product := new(utils.ProductReport)
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (dao *ReportsDaoImpl) TotalSalesByProduct() ([]*utils.TotalSold, error) {
	log.Println("Get total sales by product")
	rows, err := dao.conn.Query(`SELECT p.id, p.name, p.price, SUM(it.quantity * it.price) AS total_sold 
	FROM invoice_item it 
	INNER JOIN product p ON it.product_id = p.id 
	where it.invoice_id in (select id from invoice where status = true)
	GROUP BY p.id, p.name, p.price 
	ORDER BY total_sold DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*utils.TotalSold{}
	for rows.Next() {
		product := new(utils.TotalSold)
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.TotalSold)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// Add Paid status on invoice and price per transaction
func (dao *ReportsDaoImpl) TotalSalesByCustomers() ([]*utils.CustomerReport, error) {
	log.Println("Get total sales by customers")
	rows, err := dao.conn.Query(`
	SELECT c.id, c.name, i.status, SUM(it.quantity) AS quantity, SUM(it.quantity * it.price) AS total_sales
	FROM invoice_item it 
	INNER JOIN invoice i ON it.invoice_id = i.id 
	INNER JOIN customer c ON i.customer_id = c.id 
	INNER JOIN product p ON it.product_id = p.id
	GROUP BY c.id, c.name, i.status 
	ORDER BY i.status DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := []*utils.CustomerReport{}
	for rows.Next() {
		customer := new(utils.CustomerReport)
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Status , &customer.Quantity, &customer.TotalSales)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func (dao *ReportsDaoImpl) ReportsByStatus() ([]*utils.StatusReport, error) {
	log.Println("Get reports by status")
	rows, err := dao.conn.Query(`
	SELECT i.status, SUM(it.quantity), SUM(it.quantity * it.price) AS total_sales
	FROM invoice i
	INNER JOIN invoice_item it ON i.id = it.invoice_id
	GROUP BY i.status
	ORDER BY i.status desc
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	status := []*utils.StatusReport{}
	for rows.Next() {
		stat := new(utils.StatusReport)
		err := rows.Scan(&stat.Status, &stat.Quantity, &stat.TotalSales)
		if err != nil {
			return nil, err
		}
		status = append(status, stat)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return status, nil

}
