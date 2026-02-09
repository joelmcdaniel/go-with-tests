package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin float64

var ErrAmountNotPositive = errors.New("amount not greater than 0")
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%g BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount > 0 {
		w.balance += amount
		return nil
	} else {
		return ErrAmountNotPositive
	}
}

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
