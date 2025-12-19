package Wallet

import "fmt"

func WalletMenu() {
	myWallet := Wallet{Balance: 0, History: []Transaction{}}

	for {
		fmt.Println("\n--- Digital Wallet ---")
		fmt.Println("1. Deposit Money")
		fmt.Println("2. Spend Money")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Transaction History")
		fmt.Println("5. Back to Main Menu")
		fmt.Print("Select: ")

		var choice int
		fmt.Scan(&choice)

		if choice == 5 {
			break
		}

		switch choice {
		case 1:
			var amount float64
			var desc string
			fmt.Print("Amount: ")
			fmt.Scan(&amount)
			fmt.Print("Description: ")
			fmt.Scan(&desc)
			myWallet.AddMoney(amount, desc)
		case 2:
			var amount float64
			var desc string
			fmt.Print("Amount: ")
			fmt.Scan(&amount)
			fmt.Print("Description: ")
			fmt.Scan(&desc)
			myWallet.SpendMoney(amount, desc)
		case 3:
			fmt.Printf("Current Balance: $%.2f\n", myWallet.GetBalance())
		case 4:
			myWallet.ListTransactions()
		}
	}
}
