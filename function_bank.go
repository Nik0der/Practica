package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	owner   string
	balance float64
}

func NewBankAccount(owner string, initialBalance float64) *BankAccount {
	return &BankAccount{
		owner:   owner,
		balance: initialBalance,
	}
}

func (acc *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	acc.balance += amount
	return nil
}

func (acc *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма должна быть положительной")
	}
	if amount > acc.balance {
		return errors.New("недостаточно средств")
	}
	acc.balance -= amount
	return nil
}

func (acc BankAccount) GetBalance() float64 {
	return acc.balance
}

func main() {
	account := NewBankAccount("Алексей", 1000)

	fmt.Printf("Баланс счета %s: %.2f руб.\n", account.owner, account.GetBalance())

	err := account.Deposit(500)
	if err != nil {
		fmt.Println("Ошибка пополнения:", err)
	} else {
		fmt.Println("Счет пополнен на 500 руб.")
	}

	err = account.Withdraw(200)
	if err != nil {
		fmt.Println("Ошибка снятия:", err)
	} else {
		fmt.Println("Со счета снято 200 руб.")
	}

	err = account.Withdraw(2000)
	if err != nil {
		fmt.Println("Ошибка снятия:", err)
	}

	fmt.Printf("Итоговый баланс: %.2f руб.\n", account.GetBalance())
}
