package main

import (
	"account-transactions/internal/db"
	"account-transactions/internal/handler"
	"account-transactions/internal/repository"
	"account-transactions/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to PostgreSQL
	db.Connect()
	defer db.DB.Close()

	// Initializing repositories, services, and handlers
	accountRepo := repository.NewAccountRepository()                 // repository initialization
	accountService := service.NewAccountService(accountRepo)         // account service initialization
	transactionService := service.NewTransactionService(accountRepo) // transaction service initialization

	accountHandler := handler.NewAccountHandler(accountService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// Set up router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/accounts", accountHandler.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{account_id}", accountHandler.GetAccount).Methods("GET")
	router.HandleFunc("/transactions", transactionHandler.SubmitTransaction).Methods("POST")

	// Start the server
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
