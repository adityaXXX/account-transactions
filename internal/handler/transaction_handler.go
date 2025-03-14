package handler

import (
	"account-transactions/internal/service"
	"encoding/json"
	"net/http"
)

// TransactionHandler handles transaction-related API endpoints.
type TransactionHandler struct {
	service service.TransactionService
}

// NewTransactionHandler creates a new TransactionHandler.
func NewTransactionHandler(service service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

// SubmitTransaction handles the transaction submission API endpoint.
func (handle *TransactionHandler) SubmitTransaction(writer http.ResponseWriter, request *http.Request) {
	var transaction struct {
		SourceAccountID      uint   `json:"source_account_id"`
		DestinationAccountID uint   `json:"destination_account_id"`
		Amount               string `json:"amount"`
	}

	if err := json.NewDecoder(request.Body).Decode(&transaction); err != nil {
		http.Error(writer, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := handle.service.SubmitTransaction(transaction.SourceAccountID, transaction.DestinationAccountID, transaction.Amount); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
