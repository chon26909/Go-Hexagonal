package main

import (
	"hexagonal/handler"
	"hexagonal/repository"
	"hexagonal/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "basic:P@ssw0rd@tcp(127.0.0.1:1234)/basic")
	if err != nil {
		panic(err)
	}
	customerRepository := repository.NewCustomerRepositoryDB(db)

	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()
	router.HandleFunc("/customer", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)
	http.ListenAndServe(":4000", router)

}
