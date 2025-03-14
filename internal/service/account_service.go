package service

import (
	"account-transactions/internal/model"
	"account-transactions/internal/repository"
	"errors"
	"strconv"
)

// AccountService defines the interface for account-related operations.
type AccountService interface {
	CreateAccount(account *model.Account) error
	GetAccountByID(accountID uint) (*model.Account, error)
	UpdateAccount(account *model.Account) error
}

type accountService struct {
	accountRepo repository.AccountRepository
}

// NewAccountService returns a new instance of AccountService.
func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return &accountService{accountRepo: accountRepo}
}

// CreateAccount creates a new account.
func (accService *accountService) CreateAccount(account *model.Account) error {
	// Validate the balance to ensure it is not negative
	if balance, err := strconv.ParseFloat(account.Balance, 64); err != nil {
		return errors.New("invalid balance format")
	} else if balance < 0 {
		return errors.New("balance cannot be negative")
	}

	// Create the account in the repository
	if err := accService.accountRepo.CreateAccount(account); err != nil {
		return err
	}

	return nil
}

// GetAccountByID retrieves an account by its ID.
func (accService *accountService) GetAccountByID(accountID uint) (*model.Account, error) {
	return accService.accountRepo.GetAccountByID(accountID)
}

// UpdateAccount updates an account's information.
func (accService *accountService) UpdateAccount(account *model.Account) error {
	return accService.accountRepo.UpdateAccount(account)
}
