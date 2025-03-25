package main

import (
	"SalesDetails-CSV/config"
	"SalesDetails-CSV/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	config.InitDB()
	router := mux.NewRouter()

	router.HandleFunc("/api/revenue", handlers.GetTotalRevenue).Methods("GET")
	router.HandleFunc("/api/top-products", handlers.GetTopProducts).Methods("GET")
	router.HandleFunc("/api/customers", handlers.GetTotalCustomers).Methods("GET")
	router.HandleFunc("/api/load-data", handlers.LoadData).Methods("POST")

	fmt.Println("🚀 Server running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
