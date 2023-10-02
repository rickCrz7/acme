package main

import (
	"database/sql"
	"log"
)

type ProductDao interface {
	GetProducts() ([]*Product, error)
	GetProductById(id string) (*Product, error)
	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error
	DeleteProduct(id string) error
}

type ProductDaoImpl struct {
	conn *sql.DB
}

func NewProductDao(conn *sql.DB) *ProductDaoImpl {
	return &ProductDaoImpl{conn: conn}
}

func (dao *ProductDaoImpl) GetProducts() ([]*Product, error) {
	log.Println("Get all products")
	rows, err := dao.conn.Query("SELECT id, name, price FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*Product{}
	for rows.Next() {
		product := new(Product)
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
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

func (dao *ProductDaoImpl) GetProductById(id string) (*Product, error) {
	log.Println("Get product by id")
	row := dao.conn.QueryRow("SELECT id, name, price FROM product WHERE id = $1", id)
	product := new(Product)
	err := row.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (dao *ProductDaoImpl) CreateProduct(product *Product) error {
	log.Println("Create product")
	_, err := dao.conn.Exec("INSERT INTO product (name, price) VALUES ($1, $2)", product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func (dao *ProductDaoImpl) UpdateProduct(product *Product) error {
	log.Println("Update product")
	_, err := dao.conn.Exec("UPDATE product SET name = $1, price = $2 WHERE id = $3", product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (dao *ProductDaoImpl) DeleteProduct(id string) error {
	log.Println("Delete product")
	_, err := dao.conn.Exec("DELETE FROM product WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
