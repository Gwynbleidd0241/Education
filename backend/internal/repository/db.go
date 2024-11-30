package repository

import (
	"Kursash/internal/models"
	"Kursash/internal/storage/PostgreSQL"
	"context"
	"database/sql"
)

const users = "users"

type DB struct {
	courseStorage
	Client PostgreSQL.Client
}

type Client interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type PostgreSQLClient struct {
	DB *sql.DB
}

func (c *PostgreSQLClient) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return c.DB.ExecContext(ctx, query, args...)
}

type Authorization interface {
	CreateUser(user models.UserModel) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

func NewDB(client PostgreSQL.Client) *DB {
	return &DB{Client: client}
}
