package Employee

import "fmt"

func (p PartTime) CalculateSalary() float64 {
	return p.hourlyRate * float64(p.hoursWorked)
}

func (p PartTime) CalculateBonus() float64 {
	if p.hoursWorked > 160 {
		return 500
	}

	return 0
}

func (p PartTime) GetWorkHours() int {
	return p.hoursWorked
}

func (p PartTime) GetInfo() string {
	return fmt.Sprintf("Part-time employee: %s", p.name)
}
