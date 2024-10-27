package main

type (
	Transaction struct {
		From, To string
		Sum float64
	}

	Account struct {
		Name	string
		Balance	float64
	}
)

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

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,	
	)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	} else if (transaction.To == a.Name) {
		a.Balance += transaction.Sum
	}
	return a
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}