package Wallet

import "fmt"

func (w *Wallet) AddMoney(amount float64, desc string) {
	w.Balance += amount

	newTx := Transaction{
		Type:        "Deposit",
		Amount:      amount,
		Description: desc,
	}

	w.History = append(w.History, newTx)

	fmt.Printf("Successfully added %.2f\n", amount)
}

func (w *Wallet) SpendMoney(amount float64, desc string) {
	if amount > w.Balance {
		fmt.Println("Error: Insufficient funds")
		return
	}

	w.Balance -= amount

	newTx := Transaction{
		Type:        "Withdrawal",
		Amount:      amount,
		Description: desc,
	}

	w.History = append(w.History, newTx)

	fmt.Printf("Successfully spent %.2f\n", amount)
}

func (w *Wallet) GetBalance() float64 {
	return w.Balance
}

func (w Wallet) ListTransactions() {
	if len(w.History) == 0 {
		fmt.Println("No transactions recorded yet.")
		return
	}

	fmt.Println("--- Transactions History ---")
	fmt.Printf("%-12s | %-10s | %s\n", "Type", "Amount", "Description")
	fmt.Println("-------------------------------------------")

	for _, tx := range w.History {
		fmt.Printf("%-12s | %-10f | %s", tx.Type, tx.Amount, tx.Description)
	}
}
