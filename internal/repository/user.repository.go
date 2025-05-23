package repository

import (
	"database/sql"
	"ddb-go-demo/internal/dto"
	"ddb-go-demo/internal/schema"
	"fmt"
)

func CreateUser(db *sql.DB, user *dto.User) error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	_, err := db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func GetUserByEmail(db *sql.DB, email string) (*schema.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE email = ?"
	row := db.QueryRow(query, email)

	var user schema.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with email: %s", email)
		}
		return nil, err
	}

	return &user, nil
}
