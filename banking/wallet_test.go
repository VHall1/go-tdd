package banking

import "testing"

func TestWallet(t *testing.T) {
	w := Wallet{}
	w.Deposit(Bitcoin(10))

	got := w.Balance()
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
