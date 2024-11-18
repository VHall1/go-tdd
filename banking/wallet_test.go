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

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()

		if got == nil {
			t.Error("didn't get an error but wanted one")
		}

		if got.Error() != want {
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

		assertBalance(t, w, starting)
		assertError(t, err, "cannot withdraw, insufficient funds")
	})
}
