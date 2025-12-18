package Employee

import "fmt"

func (f FullTime) CalculateSalary() float64 {
	return f.monthlySalary
}

func (f FullTime) CalculateBonus() float64 {
	return f.monthlySalary * f.bonusRate
}

func (f FullTime) GetWorkHours() int {
	return 160
}

func (f FullTime) GetInfo() string {
	return fmt.Sprintf("Full-Time employee: %s", f.name)
}
