package Employee

type FullTime struct {
	name          string
	monthlySalary float64
	bonusRate     float64
}

type PartTime struct {
	name        string
	hourlyRate  float64
	hoursWorked int
}

type Contractor struct {
	name              string
	projectRate       float64
	projectsCompleted int
}

type Intern struct {
	name       string
	dailyRate  float64
	daysWorked int
}
