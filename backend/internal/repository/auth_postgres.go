package repository

import (
	"Kursash/internal/models"
	"database/sql"
	"fmt"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.UserModel) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password, email ) values ($1, $2, $3) RETURNING ID", users)
	row := r.db.QueryRow(query, user.Username, user.Password, user.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
