package main

import (
	"math/rand/v2"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Number   int64  `json:"number"`
	Password string `jsong:"password"`
}
type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount    int `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type Account struct {
	ID                  int       `json:"id"`
	FirstName           string    `json:"firstName"`
	LastName            string    `json:"lastName"`
	Number              int64     `json:"number"`
	EnctryptedPasssword string    `json:"-"`
	Balance             int64     `json:"balance"`
	CreatedAt           time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastname, password string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Account{
		FirstName:           firstName,
		LastName:            lastname,
		Number:              int64(rand.IntN(1000000)),
		EnctryptedPasssword: string(encpw),
		CreatedAt:           time.Now().UTC(),
	}, nil

}
