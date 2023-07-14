package accounts

// Account struct
type Account struct {
	// upper case => public
	Owner string

	// lower caee => private
	balance int
}

// NewAccount creates Account
func NewAccount(owner_ string) *Account {
	// can omit balance
	account := Account{Owner: owner_} // balance => 0
	return &account
}
