package models

import (
	"fmt"
)

type Org struct {
	Name string
	CEO  *Employee
}

func (org *Org) FindClosestCommonManager(empOne string, empTwo string) *Employee {
	result := find(org.CEO, empOne, empTwo)

	if result == nil {
		return result
	}

	if result.Name != empOne && result.Name != empTwo {
		return result
	}

	if result.Name == empOne {
		found := org.Search(result, empTwo)

		if found == nil {
			return nil
		}

		return result
	}

	if result.Name == empTwo {
		found := org.Search(result, empOne)

		if found == nil {
			return nil
		}

		return result
	}

	return result
}

func (org *Org) Search(root *Employee, searchFor string) *Employee {
	if root == nil {
		return root
	}

	if root.Name == searchFor {
		return root
	}

	for _, child := range root.Reports {
		found := org.Search(child, searchFor)

		if found != nil {
			return found
		}
	}

	return nil
}

func (org *Org) PrintLevels() {
	q := &Queue{}

	q.Enqueue(org.CEO, 1)

	prevLevel := 0

	for !q.IsEmpty() {
		item := q.Dequeue()

		if item.Level > prevLevel {
			fmt.Println()
			prevLevel = item.Level
			fmt.Print("Level ", item.Level, " - ")
		}

		fmt.Print(item.Employee.Name)
		fmt.Print(" | ")

		for _, report := range item.Employee.Reports {

			if report != nil {
				q.Enqueue(report, item.Level+1)
			}
		}
	}

	fmt.Println()
}

func find(manager *Employee, empOne string, empTwo string) *Employee {
	if manager == nil {
		return nil
	}

	if manager.Name == empOne || manager.Name == empTwo {
		return manager
	}

	count := 0
	var temp *Employee

	for _, child := range manager.Reports {
		result := find(child, empOne, empTwo)

		if result != nil {
			count++
			temp = result
		}
	}

	if count == 2 {
		return manager
	}

	return temp
}
