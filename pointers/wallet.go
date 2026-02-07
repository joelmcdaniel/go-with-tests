package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%g BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) (Bitcoin, error) {
	if amount > w.balance {
		return Bitcoin(0), ErrInsufficientFunds
	}
	w.balance -= amount
	return amount, nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
