package Gym

import "fmt"

func GymMenu() {
	myGym := Gym{Members: make(map[uint64]Member)}

	for {
		fmt.Println("\n--- Membership Management System ---")
		fmt.Println("1. Add Basic Member")
		fmt.Println("2. Add Premium Member")
		fmt.Println("3. List All Members")
		fmt.Println("4. Back to Main Menu")
		fmt.Print("Select option: ")

		var choice int
		fmt.Scan(&choice)

		if choice == 4 {
			break
		}

		switch choice {
		case 1:
			var b BasicMember
			fmt.Print("Enter ID: ")
			fmt.Scan(&b.ID)
			fmt.Print("Enter name: ")
			fmt.Scan(&b.Name)
			myGym.AddMember(b.ID, b)
		case 2:
			var p PremiumMember
			fmt.Print("Enter ID: ")
			fmt.Scan(&p.ID)
			fmt.Print("Enter name: ")
			fmt.Scan(&p.Name)
			fmt.Print("Sauna access(true/false): ")
			fmt.Scan(&p.HasSaunaAccess)
			myGym.AddMember(p.ID, p)
		case 3:
			myGym.ListMembers()
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
