package model

// Account Entity
type Account struct {
	AccountID uint   `json:"account_id" gorm:"primary_key"`
	Balance   string `json:"balance"`
}
