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

	assertError := func(t testing.TB, err , want error) {
		t.Helper()
		if err == nil {
			t.Fatal("Wanted an Error, but didn't get one")
		}

		if err != want {
			t.Errorf("Got %q, wanted %q", err, want)
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

	t.Run("Withdrawing Insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		// wallet := Wallet{balance: Bitcoin(20)}
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}