package banking

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()

		if got == nil {
			t.Error("didn't get an error but wanted one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(10))
		assertBalance(t, w, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(20)}
		w.Withdraw(Bitcoin(10))
		assertBalance(t, w, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		starting := Bitcoin(20)
		w := Wallet{starting}
		err := w.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, w, starting)
	})
}
