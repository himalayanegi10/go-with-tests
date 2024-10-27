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