package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPosrgresStorage() (*PostgresStorage, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

func (s *PostgresStorage) Init() error {
	return nil
}

func (s *PostgresStorage) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS accounts (
		id SERIAL PRIMARY KEY,
		first_name varchar(50),
		last_name varchar(50),
		account_number serial,
		balance FLOAT,
		created_at timestamp
		)`
}

func (s *PostgresStorage) CreateAccount(a *Account) error {
	return nil
}
func (s *PostgresStorage) DeleteAccount(int) error {
	return nil
}
func (s *PostgresStorage) UpdateAccount(a *Account) error {
	return nil
}
func (s *PostgresStorage) GetAccountByID(int) (*Account, error) {
	return nil, nil
}
