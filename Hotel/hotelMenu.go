package Hotel

import "fmt"

func HotelMenu() {
	h := Hotel{Rooms: make(map[string]Room)}

	for {
		fmt.Println("\n--- Hotel Management System ---")
		fmt.Println("1. Add Room")
		fmt.Println("2. Check In")
		fmt.Println("3. Check Out")
		fmt.Println("4. List Vacant Rooms")
		fmt.Println("5. Back to Main Menu")
		fmt.Print("Choice: ")

		var choice int
		fmt.Scan(&choice)

		if choice == 5 {
			break
		}

		switch choice {
		case 1:
			var r Room
			fmt.Println("Room Number: ")
			fmt.Scan(&r.RoomNumber)
			fmt.Println("Type: ")
			fmt.Scan(&r.Type)
			fmt.Println("Price(per night): ")
			fmt.Scan(&r.PricePerNight)
			h.AddRoom(r)
		case 2:
			var nr string
			fmt.Println("Enter room number to check-in: ")
			fmt.Scan(&nr)
			h.CheckIn(nr)
		case 3:
			var nr string
			fmt.Println("Enter room number to check-out: ")
			fmt.Scan(&nr)
			h.CheckOut(nr)
		case 4:
			h.ListVacantRooms()
		default:
			fmt.Println("Invalid choice")
		}
	}
}
