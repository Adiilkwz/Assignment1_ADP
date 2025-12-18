package Gym

type BasicMember struct {
	ID   uint64
	Name string
}

type PremiumMember struct {
	ID              uint64
	Name            string
	PersonalTrainer string
	HasSaunaAccess  bool
}

type Gym struct {
	Members map[uint64]Member
}
