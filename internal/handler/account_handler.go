package handler

import (
	"account-transactions/internal/model"
	"account-transactions/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// AccountHandler handles HTTP requests related to accounts.
type AccountHandler struct {
	service service.AccountService
}

// NewAccountHandler creates a new AccountHandler.
func NewAccountHandler(service service.AccountService) *AccountHandler {
	return &AccountHandler{service: service}
}

// CreateAccount handles the account creation API endpoint.
func (handle *AccountHandler) CreateAccount(writer http.ResponseWriter, request *http.Request) {
	var account model.Account
	if err := json.NewDecoder(request.Body).Decode(&account); err != nil {
		http.Error(writer, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate balance to ensure it's non-negative
	if balance, err := strconv.ParseFloat(account.Balance, 64); err != nil {
		http.Error(writer, "Invalid balance format", http.StatusBadRequest)
		return
	} else if balance < 0 {
		http.Error(writer, "Balance cannot be negative", http.StatusBadRequest)
		return
	}

	if err := handle.service.CreateAccount(&account); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
}

// GetAccount handles the account retrieval API endpoint.
func (handle *AccountHandler) GetAccount(writer http.ResponseWriter, request *http.Request) {
	accountID, err := strconv.Atoi(mux.Vars(request)["account_id"])
	if err != nil {
		http.Error(writer, "Invalid account ID", http.StatusBadRequest)
		return
	}

	account, err := handle.service.GetAccountByID(uint(accountID))
	if err != nil {
		http.Error(writer, "Account not found", http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(account)
}
