package customer

import (
	"database/sql"
	"log"

	"github.com/rickCrz7/acme/utils"
)

type CustomerDao interface {
	GetCustomers() ([]*utils.Customer, error)
	GetCustomerById(id string) (*utils.Customer, error)
	CreateCustomer(customer *utils.Customer) error
	UpdateCustomer(customer *utils.Customer) error
	DeleteCustomer(id string) error
}

type CustomerDaoImpl struct {
	conn *sql.DB
}

func NewCustomerDao(conn *sql.DB) *CustomerDaoImpl {
	return &CustomerDaoImpl{conn: conn}
}

func (dao *CustomerDaoImpl) GetCustomers() ([]*utils.Customer, error) {
	log.Println("Get all customers")
	rows, err := dao.conn.Query("SELECT id, name FROM customer ORDER BY name")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := []*utils.Customer{}
	for rows.Next() {
		customer := new(utils.Customer)
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

func (dao *CustomerDaoImpl) GetCustomerById(id string) (*utils.Customer, error) {
	log.Println("Get customer by id")
	row := dao.conn.QueryRow("SELECT id, name FROM customer WHERE id = $1", id)
	customer := new(utils.Customer)
	err := row.Scan(&customer.ID, &customer.Name)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (dao *CustomerDaoImpl) CreateCustomer(customer *utils.Customer) error {
	log.Println("Create customer")
	_, err := dao.conn.Exec("INSERT INTO customer (name) VALUES ($1)", customer.Name)
	if err != nil {
		return err
	}
	return nil
}

func (dao *CustomerDaoImpl) UpdateCustomer(customer *utils.Customer) error {
	log.Println("Update customer")
	_, err := dao.conn.Exec("UPDATE customer SET name = $1 WHERE id = $2", customer.Name, customer.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *CustomerDaoImpl) DeleteCustomer(id string) error {
	log.Println("Delete customer")
	_, err := dao.conn.Exec("DELETE FROM customer WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
