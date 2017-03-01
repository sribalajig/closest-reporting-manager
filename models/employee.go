package models

/*Employee has a name and a list of reports. The reports may be empty*/
type Employee struct {
	Name    string
	Reports []*Employee
}

/*AddReport adds a report to the employee*/
func (employee *Employee) AddReport(name string) {
	for _, report := range employee.Reports {
		if report.Name == name {
			return
		}
	}

	employee.Reports = append(employee.Reports, &Employee{Name: name})
}
