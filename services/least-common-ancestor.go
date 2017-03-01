package services

import (
	"closest-reporting-manager/models"
	"fmt"
)

func FindLeastCommonAncestor(ceo *models.Employee, empOne string, empTwo string) *models.Employee {
	result := find(ceo, empOne, empTwo)

	if result == nil {
		return result
	}

	if result.Name != empOne && result.Name != empTwo {
		return result
	}

	if result.Name == empOne {
		found := Dfs(result, empTwo)

		if found == nil {
			return nil
		}

		return result
	}

	if result.Name == empTwo {
		found := Dfs(result, empOne)

		if found == nil {
			return nil
		}

		return result
	}

	return result
}

func find(manager *models.Employee, empOne string, empTwo string) *models.Employee {
	if manager == nil {
		return nil
	}

	if manager.Name == empOne || manager.Name == empTwo {
		return manager
	}

	count := 0
	var temp *models.Employee

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

func Dfs(root *models.Employee, searchFor string) *models.Employee {
	if root == nil {
		return root
	}

	if root.Name == searchFor {
		return root
	}

	for _, child := range root.Reports {
		found := Dfs(child, searchFor)

		if found != nil {
			return found
		}
	}

	return nil
}

func Bfs(root *models.Employee) {
	q := &models.Queue{}

	q.Enqueue(root, 1)

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
