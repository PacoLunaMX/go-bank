package main

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/joho/godotenv"

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
	envFile, _ := godotenv.Read(".env")
	password := envFile["DB_PASSWORD"]

	if password == "" {
		return nil, errors.New("missing env variable DB_PASSWORD")
	}
	constStr := fmt.Sprintf("user=postgres dbname=go-bank password=%s sslmode=disable", password)

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
