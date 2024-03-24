package points

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	w.balance = w.balance + amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
