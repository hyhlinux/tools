package queue

import (
	"testing"
)

func Test(t *testing.T) {
	q := New()

	if q.Len() != 0 {
		t.Errorf("Length should be 0")
	}

	q.Enqueue(1)

	if q.Len() != 1 {
		t.Errorf("Length should be 1")
	}

	if q.Peek().(int) != 1 {
		t.Errorf("Enqueued value should be 1")
	}

	v := q.Dequeue()
	//
	t.Logf("v:%v", v)
	if v.(int) != 1 {
		t.Errorf("Dequeued value should be 1")
	}
	//
	if q.Peek() != nil || q.Dequeue() != nil {
		t.Errorf("Empty queue should have no values")
	}
	//
	t.Logf("q:%v", q)
	q.Enqueue(1)
	q.Enqueue(2)
	//
	if q.Len() != 2 {
		t.Errorf("Length should be 2")
	}
	//
	if q.Peek().(int) != 1 {
		t.Errorf("First value should be 1")
	}
	t.Logf("q:%v", q)
	//
	q.Dequeue()
	//
	t.Logf("q:%v", q)
	if q.Peek().(int) != 2 {
		t.Errorf("Next value should be 2")
	}
}
