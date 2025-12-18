package Gym

import "fmt"

func (g *Gym) AddMember(id uint64, m Member) {
	if _, exists := g.Members[id]; exists {
		fmt.Println("Error: ID already exists.")
		return
	}

	g.Members[id] = m
}

func (g *Gym) ListMembers() {
	if len(g.Members) == 0 {
		fmt.Println("No members found in the system.")
		return
	}

	fmt.Println("\n--- Current Gym Members ---")
	for _, m := range g.Members {
		fmt.Println(m.GetDetails())
	}
}
