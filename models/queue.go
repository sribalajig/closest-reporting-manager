package models

/*Node represents an item which is pushed in the queue. Level keeps track of levels in the org tree.*/
type Node struct {
	Employee *Employee
	Level    int
	Next     *Node
}

/*Queue represents a FIFO queue*/
type Queue struct {
	Front *Node
	Rear  *Node
}

/*Enqueue adds an item to the end of the queue*/
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

/*Dequeue removes an item from the front of the queue*/
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

/*IsEmpty checks if the queue is empty*/
func (q *Queue) IsEmpty() bool {
	return q.Front == nil && q.Rear == nil
}
