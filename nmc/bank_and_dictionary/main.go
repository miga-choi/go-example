package main

import (
	"fmt"
	"nmc/bank_and_dictionary/accounts"
)

func main() {
	account := accounts.Account{Owner: "owner"}
	fmt.Println(account)

	newAccount := accounts.NewAccount("newOwner")
	fmt.Println(newAccount)

	newAccount.Owner = "hi"
	fmt.Println(newAccount)
}
