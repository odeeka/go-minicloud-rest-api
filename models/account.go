package models

import (
	"errors"

	"github.com/odeeka/go-minicloud-rest-api/db"
	"github.com/odeeka/go-minicloud-rest-api/utils"
)

// Account represents a user account in the system
type Account struct {
	ID       int64  `json:"id"`       // Unique identifier
	Username string `json:"username"` // Username for login
	Password string `json:"password"` // Password (stored hashed, omit in response)
}

func GetAllAccount() ([]Account, error) {
	query := "SELECT * FROM accounts"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []Account

	for rows.Next() {
		var acc Account
		err := rows.Scan(&acc.ID, &acc.Username, &acc.Password)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, acc)
	}

	return accounts, nil
}

// Save the new account
func (acc Account) Save() error {
	query := "INSERT INTO accounts(username, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(acc.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(acc.Username, hashedPassword)

	if err != nil {
		return err
	}

	accId, err := result.LastInsertId()

	acc.ID = accId
	return err
}

// func GetAccountByUsername(username string) (*Account, error) {
// 	query := `SELECT id, username, password_hash FROM accounts WHERE username = ?`
// 	row := db.DB.QueryRow(query, username)

// 	var account Account
// 	err := row.Scan(&account.ID, &account.Username, &account.Password)
// 	if err == sql.ErrNoRows {
// 		return nil, nil
// 	} else if err != nil {
// 		return nil, err
// 	}

// 	return &account, nil
// }

func (acc *Account) ValidateCredentials() error {
	query := "SELECT id, password FROM accounts WHERE username = ?"
	row := db.DB.QueryRow(query, acc.Username)

	var retrievedPassword string
	err := row.Scan(&acc.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(acc.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
