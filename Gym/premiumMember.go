package Gym

import "fmt"

func (p PremiumMember) GetDetails() string {
	return fmt.Sprintf("ID: %d | Name: %s | Personal Trainer: %s | Has sauna access: %v",
		p.ID, p.Name, p.PersonalTrainer, p.HasSaunaAccess)
}
