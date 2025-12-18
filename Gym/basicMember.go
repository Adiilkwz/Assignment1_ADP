package Gym

import "fmt"

func (b BasicMember) GetDetails() string {
	return fmt.Sprintf("ID: %d | Name: %s", b.ID, b.Name)
}
