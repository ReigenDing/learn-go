package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
	assertError := func(t *testing.T, got error, want string) {
		if got == nil {
			// t.Errorf("want an error but got nil")
			t.Fatal("didn't get an error but want one")
		}
		if got.Error() != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		// got := wallet.Balance()
		fmt.Printf("address of balance in test %p\n", &wallet.balance)
		fmt.Printf("address of wallet in test %p\n", &wallet)
		// want := Bitcoin(10)
		// if got != want {
		// 	t.Errorf("got %v want %v", got, want)
		// }
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("WithDraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.WithDraw(Bitcoin(10))
		// got := Wallet.Balance()
		// want := Bitcoin(10)
		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{startBalance}
		err := wallet.WithDraw(Bitcoin(100))
		assertBalance(t, wallet, startBalance)
		assertError(t, err, "can not withdraw, insufficient funds")

	})

}
