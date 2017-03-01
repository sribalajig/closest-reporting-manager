package models

import (
	"testing"
)

/*Enqueue three nodes and then dequeue*/
func Test_Enqueue(t *testing.T) {
	q := Queue{}

	q.Enqueue(&Employee{Name: "elon"}, 1)
	q.Enqueue(&Employee{Name: "bill"}, 2)
	q.Enqueue(&Employee{Name: "steve"}, 2)

	if q.IsEmpty() {
		t.Errorf("Expected queue to not be empty, but was empty")
	}

	result := q.Dequeue()

	if result.Employee.Name != "elon" {
		t.Errorf("Expected '%s', but got '%s'", "elon", result.Employee.Name)
	}
}

/*Enqueue three nodes and dequeue all of them and then check if queue is empty*/
func Test_Dequeue(t *testing.T) {
	q := Queue{}

	q.Enqueue(&Employee{Name: "elon"}, 1)
	q.Enqueue(&Employee{Name: "bill"}, 2)
	q.Enqueue(&Employee{Name: "steve"}, 2)

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	result := q.Dequeue()

	if result != nil {
		t.Errorf("Expected an empty queue, but found a node")
	}
}

func Test_IsEmpty(t *testing.T) {
	q := Queue{}

	if !q.IsEmpty() {
		t.Errorf("Expected an empty queue, but found a node")
	}

	q.Enqueue(&Employee{Name: "elon"}, 1)

	if q.IsEmpty() {
		t.Errorf("Expected queue to not be empty, but was empty")
	}

	q.Dequeue()

	if !q.IsEmpty() {
		t.Errorf("Expected an empty queue, but found a node")
	}
}
