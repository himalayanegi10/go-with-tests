package pointersanderrors

type Wallet struct {
	balance float64
}

func (w *Wallet) Balance() float64 { // use pointers just like C
	return w.balance
}

func (w *Wallet) Deposit(amount float64) { // use pointers just like C
	w.balance += amount
}