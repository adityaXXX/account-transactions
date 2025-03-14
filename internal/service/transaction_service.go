package service

import (
	"account-transactions/internal/db"
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
	// Start a database transaction
	tx := db.DB.Begin()
	if tx.Error != nil {
		return errors.New("failed to begin transaction")
	}

	// Ensure the transaction is committed or rolled back
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Fetch source and destination accounts within the transaction
	sourceAccount, err := tranService.accountRepo.GetAccountByIDWithinTx(tx, sourceAccountID)
	if err != nil {
		tx.Rollback()
		return errors.New("source account not found")
	}

	destinationAccount, err := tranService.accountRepo.GetAccountByIDWithinTx(tx, destinationAccountID)
	if err != nil {
		tx.Rollback()
		return errors.New("destination account not found")
	}

	// Convert balance and amount to float64
	sourceBalance, err := strconv.ParseFloat(sourceAccount.Balance, 64)
	if err != nil {
		tx.Rollback()
		return errors.New("invalid source account balance")
	}

	transactionAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		tx.Rollback()
		return errors.New("invalid transaction amount")
	}

	// Ensure the source account has sufficient balance
	if sourceBalance < transactionAmount {
		tx.Rollback()
		return errors.New("insufficient balance")
	}

	// Deduct amount from source account
	sourceBalance -= transactionAmount
	sourceAccount.Balance = strconv.FormatFloat(sourceBalance, 'f', -1, 64)
	if err := tranService.accountRepo.UpdateAccountWithinTx(tx, sourceAccount); err != nil {
		tx.Rollback()
		return err
	}

	// Convert destination account balance to float64
	destinationBalance, err := strconv.ParseFloat(destinationAccount.Balance, 64)
	if err != nil {
		tx.Rollback()
		return errors.New("invalid destination account balance")
	}

	// Add amount to destination account
	destinationBalance += transactionAmount
	destinationAccount.Balance = strconv.FormatFloat(destinationBalance, 'f', -1, 64)
	if err := tranService.accountRepo.UpdateAccountWithinTx(tx, destinationAccount); err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return errors.New("failed to commit transaction")
	}

	return nil
}
