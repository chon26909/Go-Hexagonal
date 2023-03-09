package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {

	fmt.Println("id", id)

	customer := Customer{}
	query := "select * from customers where customer_id=?"
	err := r.db.Get(&customer, query, id)

	if err != nil {
		return nil, err
	}
	return &customer, nil

}
