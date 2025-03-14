package repository

import (
	"account-transactions/internal/db"
	"account-transactions/internal/model"
)

// AccountRepository defines the interface for account operations.
type AccountRepository interface {
	CreateAccount(account *model.Account) error
	GetAccountByID(accountID uint) (*model.Account, error)
	UpdateAccount(account *model.Account) error
}

type accountRepository struct{}

// NewAccountRepository returns a new instance of AccountRepository.
func NewAccountRepository() AccountRepository {
	return &accountRepository{}
}

// CreateAccount creates a new account in the database.
func (repository *accountRepository) CreateAccount(account *model.Account) error {
	if err := db.DB.Create(account).Error; err != nil {
		return err
	}
	return nil
}

// GetAccountByID retrieves an account by its ID.
func (repository *accountRepository) GetAccountByID(accountID uint) (*model.Account, error) {
	var account model.Account
	if err := db.DB.First(&account, accountID).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

// UpdateAccount updates the balance of an existing account.
func (repository *accountRepository) UpdateAccount(account *model.Account) error {
	if err := db.DB.Save(account).Error; err != nil {
		return err
	}
	return nil
}
