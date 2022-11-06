package main

import (
	"errors"
	"fmt"
)

var ErrorInsuficientBalance = errors.New("No is possible withdraw: insuficient balance")

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(quantity Bitcoin) {
	w.balance += quantity
}

func (w *Wallet) Withdrawal(quantity Bitcoin) error {
	if quantity > w.balance {
		return ErrorInsuficientBalance
	}

	w.balance -= quantity

	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
