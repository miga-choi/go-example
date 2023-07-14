package main

import (
	"fmt"
	"log"
	"nmc/bank_and_dictionary/accounts"
)

func main() {
	account := accounts.Account{Owner: "owner"}
	fmt.Println(account)

	newAccount := accounts.NewAccount("newOwner")
	fmt.Println(newAccount)

	newAccount.Owner = "hi"
	fmt.Println(newAccount)

	account.WrongDeposit(100)
	fmt.Println(account.Balance())

	account.Deposit(100)
	fmt.Println(account.Balance())

	err := account.Withdraw(10)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(account.Balance())
	}

	err = account.Withdraw(100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(account.Balance())
	}
}
