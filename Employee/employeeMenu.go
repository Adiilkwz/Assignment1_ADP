package Employee

import "fmt"

func EmployeeMenu() {
	var employees []Employee

	for {
		fmt.Println("\n--- Employee Management ---")
		fmt.Println("1. Add Full-Time")
		fmt.Println("2. Add Part-Time")
		fmt.Println("3. Add Contractor")
		fmt.Println("4. Add Intern")
		fmt.Println("5. List All Employees & Salaries")
		fmt.Println("6. Back to Main Menu")
		fmt.Print("Choice: ")

		var choice int
		fmt.Scan(&choice)

		if choice == 6 {
			break
		}

		switch choice {
		case 1:
			var f FullTime
			fmt.Print("Name: ")
			fmt.Scan(&f.name)
			fmt.Print("Monthly Salary: ")
			fmt.Scan(&f.monthlySalary)
			fmt.Print("Bonus rate(e.g 0.1): ")
			fmt.Scan(&f.bonusRate)
			employees = append(employees, f)
		case 2:
			var p PartTime
			fmt.Print("Name: ")
			fmt.Scan(&p.name)
			fmt.Print("Hourly Rate: ")
			fmt.Scan(&p.hourlyRate)
			fmt.Print("Hours worked: ")
			fmt.Scan(&p.hoursWorked)
			employees = append(employees, p)
		case 3:
			var c Contractor
			fmt.Print("Name: ")
			fmt.Scan(&c.name)
			fmt.Print("Project rate: ")
			fmt.Scan(&c.projectRate)
			fmt.Print("Project completed: ")
			fmt.Scan(&c.projectsCompleted)
			employees = append(employees, c)
		case 4:
			var i Intern
			fmt.Print("Name: ")
			fmt.Scan(&i.name)
			fmt.Print("Daily rate: ")
			fmt.Scan(&i.dailyRate)
			fmt.Print("Days worked: ")
			fmt.Scan(&i.daysWorked)
			employees = append(employees, i)
		case 5:
			if len(employees) == 0 {
				fmt.Println("Employee list is empty.")
				break
			}

			fmt.Println("--- Employee Payroll ---")
			for _, e := range employees {
				salary := e.CalculateSalary()
				bonus := e.CalculateBonus()
				fmt.Printf("%s | Work: %d units | Salary: %.2f | Bonus: %.2f | Total: %.2f\n",
					e.GetInfo(), e.GetWorkHours(), salary, bonus, salary+bonus)
			}
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
