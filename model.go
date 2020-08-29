//RESTFULL SERVICES

package main

import (
	"database/sql"
	"errors"
)

type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

func (p *Product) getProduct(db *sql.DB) error {

	return errors.New("Not implemented")
}
func (p *Product) addProduct(db *sql.DB) error {

	return errors.New("Not implemented")
}
func (p *Product) updateProduct(db *sql.DB) error {

	return errors.New("Not implemented")
}

func getproducts(db *sql.DB, start, count int) ([]Product, error) {

	return nil, errors.New("not implemented")
}
