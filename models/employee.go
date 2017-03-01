package models

type Employee struct {
	Name    string
	Reports []*Employee
}

func (employee *Employee) AddReport(name string) {
	for _, report := range employee.Reports {
		if report.Name == name {
			return
		}
	}

	employee.Reports = append(employee.Reports, &Employee{Name: name})
}
