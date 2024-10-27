package main

type Transaction struct {
	From, To string
	Sum float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	var balance float64
	for _, transaction := range transactions {
		if transaction.From == name {
			balance -= transaction.Sum
		} else if (transaction.To == name) {
			balance += transaction.Sum
		}
	}
	return balance
}