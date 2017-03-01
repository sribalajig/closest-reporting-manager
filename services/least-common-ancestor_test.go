package services

import (
	"closest-reporting-manager/models"
	"testing"
)

/*
The following tests are based on this company tree with elon musk as the ceo

							elon musk
						/		 |		\
			mark zukerberg	bill gates 	larry paige
					|				|			\
			    steve jobs		sergei brin		sachin bansal
			/		|		\
satya nadella  paul allen	marvin minsky
*/
var companyTree = &models.Employee{
	Name: "elon musk",
	Reports: []*models.Employee{
		&models.Employee{
			Name: "mark zukerberg",
			Reports: []*models.Employee{
				&models.Employee{
					Name: "steve jobs",
					Reports: []*models.Employee{
						&models.Employee{
							Name: "satya nadella",
						},
						&models.Employee{
							Name: "paul allen",
						},
						&models.Employee{
							Name: "marvin minsky",
						},
					},
				},
			},
		},
		&models.Employee{
			Name: "bill gates",
			Reports: []*models.Employee{
				&models.Employee{
					Name: "sergei brin",
				},
			},
		},
		&models.Employee{
			Name: "larry paige",
			Reports: []*models.Employee{
				&models.Employee{
					Name: "sachin bansal",
				},
			},
		},
	},
}

var testCases = []struct {
	ceo            *models.Employee
	empOne         string
	empTwo         string
	expectedResult string
}{
	{
		companyTree,
		"elon musk",
		"bill gates",
		"elon musk",
	},
	{
		companyTree,
		"steve jobs",
		"bill gates",
		"elon musk",
	},
	{
		companyTree,
		"larry paige",
		"sachin bansal",
		"larry paige",
	},
	{
		companyTree,
		"paul allen",
		"marvin minsky",
		"steve jobs",
	},
	{
		companyTree,
		"bill gates",
		"marvin minsky",
		"elon musk",
	},
	{
		companyTree,
		"ratan tata",   // not in the company
		"binny bansal", // not in the company
		"",             // there should be no common manager
	},
	{
		companyTree,
		"sachin bansal", // in the company
		"ratan tata",    // not in the company
		"",              // there should be no common manager
	},
	{
		companyTree,
		"elon musk",  // in the company
		"katy perry", // not in the company
		"",           // there should be no common manager
	},
}

/*Runs the least common ancestor algorithm across different test cases*/
func Test_LeastCommonAncestor(t *testing.T) {
	for _, testCase := range testCases {
		lca := FindLeastCommonAncestor(testCase.ceo, testCase.empOne, testCase.empTwo)

		if lca == nil && testCase.expectedResult != "" {
			t.Fail()

			t.Errorf("Expected not nil, got nil for %s, %s", testCase.empOne, testCase.empTwo)
		} else if lca != nil && lca.Name != testCase.expectedResult {
			t.Fail()

			t.Errorf("Expected %s, Got %s", testCase.expectedResult, lca.Name)
		}
	}
}

/*This test prints the org by levels, run the test to see the magic :)*/
func Test_BFS(t *testing.T) {
	Bfs(companyTree)
}
