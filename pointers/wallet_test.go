package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		t.Run("deposit amount positive", func(t *testing.T) {
			wallet := Wallet{}
			err := wallet.Deposit(Bitcoin(.05))
			assertNoError(t, err)
			assertBalance(t, wallet, Bitcoin(.05))

		})

		t.Run("deposit amount not positive", func(t *testing.T) {
			wallet := Wallet{}
			err := wallet.Deposit(Bitcoin(0.0))
			assertError(t, err, ErrAmountNotPositive)
			assertBalance(t, wallet, wallet.Balance())

		})
	})

	t.Run("withdraw", func(t *testing.T) {
		t.Run("withdraw with funds", func(t *testing.T) {
			wallet := Wallet{balance: Bitcoin(1.05)}
			amount, err := wallet.Withdraw(Bitcoin(.5))
			assertNoError(t, err)
			assertBalance(t, wallet, Bitcoin(.55))
			assertWithdrawnAmount(t, amount, Bitcoin(.5))
		})

		t.Run("withdraw without funds", func(t *testing.T) {
			wallet := Wallet{balance: Bitcoin(1.05)}
			amount, err := wallet.Withdraw(Bitcoin(5))
			assertError(t, err, ErrInsufficientFunds)
			assertBalance(t, wallet, Bitcoin(1.05))
			assertWithdrawnAmount(t, amount, Bitcoin(0))
		})
	})

}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertWithdrawnAmount(t testing.TB, got Bitcoin, want Bitcoin) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
