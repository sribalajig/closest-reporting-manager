package models

type Node struct {
	Employee *Employee
	Level    int
	Next     *Node
}

type Queue struct {
	Front *Node
	Rear  *Node
}

func (q *Queue) Enqueue(Employee *Employee, Level int) {
	node := &Node{
		Employee: Employee,
		Level:    Level,
	}

	if q.Front == nil && q.Front == q.Rear {
		q.Front = node
		q.Rear = node

		return
	}

	q.Rear.Next = node
	q.Rear = node
}

func (q *Queue) Dequeue() *Node {
	if q.Front == q.Rear {
		temp := q.Front

		q.Front = nil
		q.Rear = nil

		return temp
	}

	temp := q.Front
	q.Front = temp.Next

	if q.Front == nil {
		q.Rear = nil
	}

	return temp
}

func (q *Queue) IsEmpty() bool {
	return q.Front == nil && q.Rear == nil
}
