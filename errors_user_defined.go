package main

import "fmt"

type InsufficientBalanceError struct {
	Balance int
	Amount  int
}

func (e *InsufficientBalanceError) Error() string {
	return fmt.Sprintf("insufficient balance: balance is %d, attempted to withdraw %d", e.Balance, e.Amount)
}

func withdraw(balance, amount int) error {
	if amount > balance {
		return &InsufficientBalanceError{Balance: balance, Amount: amount}
	}
	return nil
}

func main() {
	err := withdraw(1000, 1500)

	if err != nil {
		fmt.Println("Error:", err)
	}