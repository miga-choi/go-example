package accounts

import "errors"

// Account struct
type Account struct {
	// upper case => public
	Owner string

	// lower caee => private
	balance int
}

var errInsufficientBalance = errors.New("Insufficient balance!")

// NewAccount creates Account
func NewAccount(owner_ string) *Account {
	// can omit balance
	account := Account{Owner: owner_} // balance => 0
	return &account
}

// Deposit amount on your account
// account_ without "*", account_ is a copy of parameter, so origin account_ won't increase
func (account_ Account) WrongDeposit(amount_ int) {
	account_.balance += amount_
}

// account_ with "*", account_ point to origin account_, so origin account_ will increase
func (account_ *Account) Deposit(amount_ int) {
	account_.balance += amount_
}

// Withdraw amount from your account
func (account_ *Account) Withdraw(amount_ int) error {
	if account_.balance < amount_ {
		return errInsufficientBalance
	}
	account_.balance -= amount_
	return nil
}

// Balance of your account
func (account_ Account) Balance() int {
	return account_.balance
}
