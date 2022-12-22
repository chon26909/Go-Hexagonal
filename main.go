package main

import (
	"fmt"
	"hexagonal/repository"
	"hexagonal/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:password@tcp(127.0.0.1:3306)/basic")
	if err != nil {
		panic(err)
	}
	customerRepository := repository.NewCustomerRepositoryDB(db)

	customerService := service.NewCustomerService(customerRepository)

	// customers, err := customerService.GetCustomers()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("data", customers)

	customer, err := customerService.GetCustomer(0)
	if err != nil {
		panic(err)
	}

	fmt.Println("customer", customer)

}
