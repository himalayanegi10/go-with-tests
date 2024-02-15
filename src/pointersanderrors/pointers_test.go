package pointersanderrors

import (
	"testing"
)

func TestPointers(t *testing.T) {

	assertBalance := func(t testing.TB, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()
		if got != want {
			t.Errorf("Expected %s, but got %s", want, got)
		}
	}

	t.Run("Bitcoin Wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Wallet Withdraw Bitcoin", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)}
		wallet.Withdraw(Bitcoin(14))
		assertBalance(t, wallet, Bitcoin(16))	
	})
}