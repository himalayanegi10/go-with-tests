package main

type Transaction struct {
	From, To string
	Sum float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			currentBalance -= t.Sum
		} else if (t.To == name) {
			currentBalance += t.Sum
		}
		return currentBalance
	}
	return Reduce(transactions, adjustBalance, 0.0)
}