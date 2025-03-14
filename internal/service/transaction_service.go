package service

import (
	"account-transactions/internal/repository"
	"errors"
	"strconv"
)

// TransactionService defines the interface for transaction-related operations.
type TransactionService interface {
	SubmitTransaction(sourceAccountID uint, destinationAccountID uint, amount string) error
}

type transactionService struct {
	accountRepo repository.AccountRepository
}

// NewTransactionService returns a new instance of TransactionService.
func NewTransactionService(accountRepo repository.AccountRepository) TransactionService {
	return &transactionService{accountRepo: accountRepo}
}

// SubmitTransaction processes a transaction between two accounts.
func (tranService *transactionService) SubmitTransaction(sourceAccountID uint, destinationAccountID uint, amount string) error {
	// Fetch source and destination accounts
	sourceAccount, err := tranService.accountRepo.GetAccountByID(sourceAccountID)
	if err != nil {
		return errors.New("source account not found")
	}

	destinationAccount, err := tranService.accountRepo.GetAccountByID(destinationAccountID)
	if err != nil {
		return errors.New("destination account not found")
	}

	// Convert balance and amount to float64
	sourceBalance, err := strconv.ParseFloat(sourceAccount.Balance, 64)
	if err != nil {
		return errors.New("invalid source account balance")
	}

	transactionAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return errors.New("invalid transaction amount")
	}

	// Ensure the source account has sufficient balance
	if sourceBalance < transactionAmount {
		return errors.New("insufficient balance")
	}

	// Deduct amount from source account
	sourceBalance -= transactionAmount
	sourceAccount.Balance = strconv.FormatFloat(sourceBalance, 'f', -1, 64)
	if err := tranService.accountRepo.UpdateAccount(sourceAccount); err != nil {
		return err
	}

	// Convert destination account balance to float64
	destinationBalance, err := strconv.ParseFloat(destinationAccount.Balance, 64)
	if err != nil {
		return errors.New("invalid destination account balance")
	}

	// Add amount to destination account
	destinationBalance += transactionAmount
	destinationAccount.Balance = strconv.FormatFloat(destinationBalance, 'f', -1, 64)
	if err := tranService.accountRepo.UpdateAccount(destinationAccount); err != nil {
		return err
	}

	return nil
}
