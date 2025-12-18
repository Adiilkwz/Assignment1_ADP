package Employee

import "fmt"

func (i Intern) CalculateSalary() float64 {
	return i.dailyRate * float64(i.daysWorked)
}

func (i Intern) CalculateBonus() float64 {
	if i.daysWorked > 20 {
		return 200.0
	}

	return 0
}

func (i Intern) GetWorkHours() int {
	return i.daysWorked * 8
}

func (i Intern) GetInfo() string {
	return fmt.Sprintf("Intern: %s", i.name)
}
