package main

import "math/rand/v2"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstName, lastname string) *Account {
	return &Account{
		ID:        int(rand.IntN(1000)),
		FirstName: firstName,
		LastName:  lastname,
		Number:    int64(rand.IntN(1000000)),
	}

}
