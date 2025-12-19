package Wallet

type Transaction struct {
	Type        string
	Amount      float64
	Description string
}

type Wallet struct {
	Balance float64
	History []Transaction
}
