package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p\n", &w.balance)
	fmt.Printf("address of wallet in Deposit is %p\n", w)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

type String interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) WithDraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("can not withdraw, insufficient funds")
	}
	w.balance -= amount
	return nil
}
