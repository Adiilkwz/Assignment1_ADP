package Employee

import "fmt"

func (c Contractor) CalculateSalary() float64 {
	return c.projectRate * float64(c.projectsCompleted)
}

func (c Contractor) CalculateBonus() float64 {
	return float64(c.projectsCompleted) * 100.0
}

func (c Contractor) GetWorkHours() int {
	return c.projectsCompleted * 40
}

func (c Contractor) GetInfo() string {
	return fmt.Sprintf("Contractor employee: %s", c.name)
}
