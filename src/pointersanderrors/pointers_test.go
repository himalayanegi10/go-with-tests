package pointersanderrors

import (
	"testing"
)

func TestPointers(t *testing.T) {
	t.Run("Bitcoin Wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.0)
		
		got := wallet.Balance()
		
		want := 10.0

		if got != want {
			t.Errorf("Wanted %g got %g", want, got)
		}
	})
}