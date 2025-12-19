package main

import (
	"fmt"

	"github.com/AdilRamazan/Assignment1/Employee"
	"github.com/AdilRamazan/Assignment1/Gym"
	"github.com/AdilRamazan/Assignment1/Hotel"
	"github.com/AdilRamazan/Assignment1/Wallet"
)

func main() {
	for {
		fmt.Println("\n--- Choose the System ---")
		fmt.Println("1. Hotel Room Reservation System")
		fmt.Println("2. Employee Management System")
		fmt.Println("3. Membership Management System")
		fmt.Println("4. Wallet Transaction Simulation")
		fmt.Println("5. Exit")
		fmt.Print("Select: ")

		var choice int
		fmt.Scan(&choice)

		if choice == 5 {
			break
		}

		switch choice {
		case 1:
			Hotel.HotelMenu()
		case 2:
			Employee.EmployeeMenu()
		case 3:
			Gym.GymMenu()
		case 4:
			Wallet.WalletMenu()
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
