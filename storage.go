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

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	constStr := "user=postgres dbname=go-bank password=Herrpep1*1996 sslmode=disable"
	db, err := sql.Open("postgres", constStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) CreateAccount(*Account) error {
	return nil

}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil

}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil

}
func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil

}
