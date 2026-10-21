package main

import (
    "errors"
    "fmt"
)

func withdraw(balance, amount int) error {
    if amount > balance {
        return errors.New("insufficient balance")
    }
    return nil
}

func main() {
    err := withdraw(1000, 1500)

    if err != nil {
        fmt.Println("Error:", err)
    }
}