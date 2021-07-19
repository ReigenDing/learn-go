package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

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
		got := wallet.WithDraw(Bitcoin(10))
		// got := Wallet.Balance()
		// want := Bitcoin(10)
		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, got)
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startBalance := Bitcoin(20)
		wallet := Wallet{startBalance}
		err := wallet.WithDraw(Bitcoin(100))
		assertBalance(t, wallet, startBalance)
		assertError(t, err, InsufficientFundsError)

	})

}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("did not get error but want one")
	}
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but did not want one")
	}
}
