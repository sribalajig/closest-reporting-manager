package models

import (
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
var ceo = &Employee{
	Name: "elon musk",
	Reports: []*Employee{
		&Employee{
			Name: "mark zukerberg",
			Reports: []*Employee{
				&Employee{
					Name: "steve jobs",
					Reports: []*Employee{
						&Employee{
							Name: "satya nadella",
						},
						&Employee{
							Name: "paul allen",
						},
						&Employee{
							Name: "marvin minsky",
						},
					},
				},
			},
		},
		&Employee{
			Name: "bill gates",
			Reports: []*Employee{
				&Employee{
					Name: "sergei brin",
				},
			},
		},
		&Employee{
			Name: "larry paige",
			Reports: []*Employee{
				&Employee{
					Name: "sachin bansal",
				},
			},
		},
	},
}

var org = Org{
	CEO: ceo,
}

var testCases = []struct {
	empOne         string
	empTwo         string
	expectedResult string
}{
	{
		"elon musk",
		"bill gates",
		"elon musk",
	},
	{
		"steve jobs",
		"bill gates",
		"elon musk",
	},
	{
		"larry paige",
		"sachin bansal",
		"larry paige",
	},
	{
		"paul allen",
		"marvin minsky",
		"steve jobs",
	},
	{
		"bill gates",
		"marvin minsky",
		"elon musk",
	},
	{
		"ratan tata",   // not in the company
		"binny bansal", // not in the company
		"",             // there should be no common manager
	},
	{
		"sachin bansal", // in the company
		"ratan tata",    // not in the company
		"",              // there should be no common manager
	},
	{
		"elon musk",  // in the company
		"katy perry", // not in the company
		"",           // there should be no common manager
	},
}

/*Runs the least common ancestor algorithm across different test cases*/
func Test_LeastCommonAncestor(t *testing.T) {
	for _, testCase := range testCases {
		lca := org.FindClosestCommonManager(testCase.empOne, testCase.empTwo)

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
	org.PrintLevels()
}
