package pointersanderrors

import (
	"testing"
)

func TestPointers(t *testing.T) {
	t.Run("Bitcoin Wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))
		
		got := wallet.Balance()
		
		want := Bitcoin(10.0)

		if got != want {
			t.Errorf("Wanted %s got %s", want, got)
		}
	})
}