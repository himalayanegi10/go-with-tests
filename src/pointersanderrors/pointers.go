package pointersanderrors

import "fmt"

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Balance() Bitcoin { // use pointers just like C
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) { // use pointers just like C
	w.balance += amount
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%g BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}