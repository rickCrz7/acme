package main

import (
	"database/sql"
	"log"
)

type CustomerDao interface {
	GetCustomers() ([]*Customer, error)
	GetCustomerById(id int) (*Customer, error)
	CreateCustomer(customer *Customer) error
	UpdateCustomer(customer *Customer) error
	DeleteCustomer(id int) error
}

type CustomerDaoImpl struct {
	conn *sql.DB
}

func NewCustomerDao(conn *sql.DB) *CustomerDaoImpl {
	return &CustomerDaoImpl{conn: conn}
}

func (dao *CustomerDaoImpl) GetCustomers() ([]*Customer, error) {
	log.Println("Get all customers")
	rows, err := dao.conn.Query("SELECT id, name FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := []*Customer{}
	for rows.Next() {
		customer := new(Customer)
		err := rows.Scan(&customer.ID, &customer.Name)
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

func (dao *CustomerDaoImpl) GetCustomerById(id int) (*Customer, error) {
	log.Println("Get customer by id")
	row := dao.conn.QueryRow("SELECT id, name FROM customers WHERE id = $1", id)
	customer := new(Customer)
	err := row.Scan(&customer.ID, &customer.Name)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (dao *CustomerDaoImpl) CreateCustomer(customer *Customer) error {
	log.Println("Create customer")
	_, err := dao.conn.Exec("INSERT INTO customers (name) VALUES ($1)", customer.Name)
	if err != nil {
		return err
	}
	return nil
}

func (dao *CustomerDaoImpl) UpdateCustomer(customer *Customer) error {
	log.Println("Update customer")
	_, err := dao.conn.Exec("UPDATE customers SET name = $1 WHERE id = $2", customer.Name, customer.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *CustomerDaoImpl) DeleteCustomer(id int) error {
	log.Println("Delete customer")
	_, err := dao.conn.Exec("DELETE FROM customers WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}