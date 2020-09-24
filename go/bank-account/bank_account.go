package account

import "sync"

// Account represents a bank account
type Account struct {
	balance int64
	closed  bool
	mux     sync.Mutex
}

// Open creates an account with an initial deposit
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{balance: initialDeposit}
}

// Close closes the account
func (a *Account) Close() (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.closed {
		return 0, false
	}

	payout := a.balance
	a.balance = 0
	a.closed = true
	return payout, true
}

// Balance gives the current balance of the account
func (a *Account) Balance() (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.closed {
		return 0, false
	}

	return a.balance, true
}

// Deposit adds/removes ammount from the account
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.closed {
		return 0, false
	}

	if a.balance+amount < 0 {
		return 0, false
	}

	a.balance += amount
	return a.balance, true
}
