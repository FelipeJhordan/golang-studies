package main

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		expect := Bitcoin(10)

		confirmBalance(t, wallet, expect)
	})

	t.Run("Withdrawal", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}

		wallet.Withdrawal(Bitcoin(10))

		expect := Bitcoin(10)

		confirmBalance(t, wallet, expect)
	})

	t.Run("Withdrawal with insufficiente value", func(t *testing.T) {
		initialBalance := Bitcoin(20)

		wallet := Wallet{initialBalance}

		erro := wallet.Withdrawal(Bitcoin(100))

		confirmBalance(t, wallet, initialBalance)
		confirmError(t, erro, ErrorInsuficientBalance)
	})
}

func confirmBalance(t *testing.T, wallet Wallet, expect Bitcoin) {
	t.Helper()

	result := wallet.Balance()

	if result != expect {
		t.Errorf("result %s, expect %s", result, expect)
	}
}

func confirmError(t *testing.T, errorResult error, expect error) {
	t.Helper()

	if errorResult == nil {
		t.Errorf("expect an error, but any occurred")
	}

	if errorResult != expect {
		t.Errorf("result %s, expect %s", errorResult, expect)
	}
}
