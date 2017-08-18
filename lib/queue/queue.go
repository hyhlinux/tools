package queue

type (
	// Queue FIFO
	Queue struct {
		start, end *node
		length     int
	}
	node struct {
		value interface{}
		next  *node
	}
)

// New Create a new queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

// Len Return the number of items in the queue
func (q *Queue) Len() int {
	return q.length
}

// Enqueue Put an item on the end of a queue
func (q *Queue) Enqueue(value interface{}) {
	n := &node{value, nil}
	if q.length != 0 {
		q.end.next = n
		q.end = n
	} else {
		q.start, q.end = n, n
	}
	q.length++
}

// Dequeue Take the next item off the front of the queue
func (q *Queue) Dequeue() interface{} {
	if q.length == 0 {
		return nil
	}

	n := q.start
	if q.length != 1 {
		q.start = q.start.next
	} else {
		q.start = nil
		q.end = nil
	}

	q.length--
	return n.value
}

// Peek Return the first item in the queue without removing it
func (q *Queue) Peek() interface{} {
	if q.length == 0 {
		return nil
	}
	return q.start.value
}
