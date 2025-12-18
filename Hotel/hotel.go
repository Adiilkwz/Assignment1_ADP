package Hotel

import (
	"fmt"
)

func (h *Hotel) AddRoom(r Room) {
	_, exists := h.Rooms[r.RoomNumber]

	if exists {
		fmt.Println("Room is already exists.")
		return
	}

	h.Rooms[r.RoomNumber] = r
	fmt.Printf("Room %s was added succesfully.\n", r.RoomNumber)
}

func (h *Hotel) CheckIn(roomNumber string) {
	room, exists := h.Rooms[roomNumber]

	if !exists {
		fmt.Println("Room does not exist.")
		return
	}

	if room.IsOccupied {
		fmt.Println("Room is already occupied.")
		return
	} else {
		room.IsOccupied = true
		h.Rooms[roomNumber] = room
		fmt.Printf("Succesfully checked-in to %s room.\n", roomNumber)
	}
}

func (h *Hotel) CheckOut(roomNumber string) {
	room, exists := h.Rooms[roomNumber]

	if !exists {
		fmt.Println("Room does not exist.")
		return
	}
	if room.IsOccupied {
		room.IsOccupied = false
		h.Rooms[roomNumber] = room
		fmt.Printf("Succesfully checked out from %s room.\n", roomNumber)
	} else {
		fmt.Println("Room is already vacant.")
	}
}

func (h *Hotel) ListVacantRooms() {
	fmt.Println("Vacant Rooms:")
	for _, room := range h.Rooms {
		if !room.IsOccupied {
			fmt.Printf("No: %s | Type: %s | Price(per night): %.2f\n", room.RoomNumber, room.Type, room.PricePerNight)
		}
	}
}
