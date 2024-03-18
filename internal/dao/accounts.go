package dao

import (
	"database/sql"
)

type AccountsRepository struct {
	DB *sql.DB
}

type Account struct {
	Username string
	Password string
	Role     int
}

type AccountsRepositoryInterface interface {
	GetAccount(username string) (*Account, error)
}

func (accr AccountsRepository) GetAccount(username string) (*Account, error) {
	row := accr.DB.QueryRow("SELECT username, password, role FROM accounts WHERE username = $1", username)

	var account Account
	err := row.Scan(&account.Username, &account.Password, &account.Role)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
