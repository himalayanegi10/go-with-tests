package main

import (
	"testing"
)

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "A",
			To: "B",
			Sum: 100,
		},
		{
			From: "C",
			To: "A",
			Sum: 25,
		},
	}

	AssertEqual(t, BalanceFor(transactions, "B"), 100)
	AssertEqual(t, BalanceFor(transactions, "A"), -75)
	AssertEqual(t, BalanceFor(transactions, "C"), -25)
}

func TestNewBadBank(t *testing.T) {
	var (
		himalaya = Account{Name: "Himalaya", Balance: 100}
		mahendra = Account{Name: "Mahendra", Balance: 75}
		tulsi = Account{Name: "Tulsi", Balance: 200}

		transactions = []Transaction{
			NewTransaction(himalaya, mahendra, 100),
			NewTransaction(himalaya, tulsi, 500),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(himalaya), -500)
	AssertEqual(t, newBalanceFor(mahendra), 175)
	AssertEqual(t, newBalanceFor(tulsi), 700)
}