package Employee

type Employee interface {
	CalculateSalary() float64
	GetWorkHours() int
	CalculateBonus() float64
	GetInfo() string
}
