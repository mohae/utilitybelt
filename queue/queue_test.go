package queue

import (
	"testing"
)

func TestNew(t *testing.T) {
	q := New(10, 0)
	if cap(q.items) != 10 {
		t.Errorf("expected 10, got %d", cap(q.items))
	}
	if q.maxCap != 0 {
		t.Errorf("expected 0, got %d", q.maxCap)
	}

	q = New(100, 200)
	if cap(q.items) != 100 {
		t.Errorf("expected 100, got %d", cap(q.items))
	}
	if q.maxCap != 200 {
		t.Errorf("expected 200, got %d", q.maxCap)
	}
}

// tests enqueue, growth, capacity restriction, and basic dequeue
func TestQueueing(t *testing.T) {
	var tests = []struct {
		size        int
		maxCap      int
		headPos     int
		tailPos     int
		expectedCap int
		items       []interface{}
		errString   string
	}{
		{size: 2, maxCap: 0, tailPos: 4, expectedCap: 4, items: []interface{}{0, 1, 2, 3}, errString: ""},
		{size: 2, maxCap: 0, tailPos: 5, expectedCap: 8, items: []interface{}{0, 1, 2, 3, 4}, errString: ""},
		{size: 2, maxCap: 4, tailPos: 4, expectedCap: 4, items: []interface{}{0, 1, 2, 3, 4}, errString: "groweQueue: cannot grow beyond max capacity of 4"},
	}
	for _, test := range tests {
		var err error
		q := New(test.size, test.maxCap)
		for _, v := range test.items {
			err = q.Enqueue(v)
		}
		if test.errString != "" {
			if err == nil {
				t.Errorf("Expected error, got none")
				continue
			}
			if err.Error() != test.errString {
				t.Errorf("Expected error to be %q. got %q", test.errString, err.Error())
				continue
			}
		}
		// check that the items are as expected:
		if len(q.items) != test.expectedCap {
			t.Error("Expected %d items in queue, got %d", test.expectedCap, len(q.items))
		}
		if cap(q.items) != test.expectedCap {
			t.Error("Expected queue cap to be %d, got %d", test.expectedCap, cap(q.items))
		}
		if q.head != test.headPos {
			t.Errorf("Expected head to be at pos %d, got %d", test.headPos, q.head)
		}
		if q.maxCap != test.maxCap {
			t.Errorf("Expected maxCap to be %d, was %d", test.maxCap, q.maxCap)
		}
		if cap(q.items) > test.maxCap && test.maxCap > 0 {
			t.Errorf("Expected cap of queue to be equal to it's max capacity, %d; was %d", test.maxCap, cap(q.items))
		}
		for i := 0; i < q.tail; i++ {
			if q.items[i] != test.items[i] {
				t.Errorf("Expected value of index %d to be %d, got %d", i, test.items[i], q.items[i])
			}
		}

		// dequeue 1 item and check
		next := q.Dequeue()
		if next != test.items[0] {
			t.Errorf("Expected %d, got %d", test.items[0], next)
			continue
		}
		if q.head != 1 {
			t.Errorf("Expected head to point to 1, got %d", q.head)
		}
	}
}

// Tests Enqueue/Dequeue/Enqueue, shifting, and growth is properly handled
func TestDequeueEngueue(t *testing.T) {
	tests := []struct {
		size        int
		maxCap      int
		headPos     int
		tailPos     int
		expectedCap int
		dequeueCnt  int
		dequeueVals []interface{}
		items       []interface{}
		items2      []interface{}
		errString   string
	}{
		{size: 10, maxCap: 0, tailPos: 9, expectedCap: 10, dequeueCnt: 3, dequeueVals: []interface{}{0, 1, 2},
			items: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, items2: []interface{}{10, 11}, errString: ""},
	}

	// First add the queue
	for _, test := range tests {
		var err error
		q := New(test.size, test.maxCap)
		for _, v := range test.items {
			err = q.Enqueue(v)
		}
		if test.errString != "" {
			if err == nil {
				t.Errorf("Expected error, got none")
				continue
			}
			if err.Error() != test.errString {
				t.Errorf("Expected error to be %q. got %q", test.errString, err.Error())
				continue
			}
		}

		// dequeue 3 items
		for i := 0; i < test.dequeueCnt; i++ {
			v := q.Dequeue()
			if v != test.dequeueVals[i] {
				t.Errorf("Expected %v, got %v", test.dequeueVals[i], v)
			}
		}

		if q.head != test.dequeueCnt {
			t.Errorf("Expected head to point to %d, got %d", test.dequeueCnt, q.head)
		}
		// enqueue the next items; should not grow, should just shift the items
		for _, v := range test.items2 {
			q.Enqueue(v)
		}
		if q.head != 0 {
			t.Errorf("Expected head to be at pos 0, got %d", q.head)
		}
		if q.tail != test.tailPos {
			t.Errorf("Expected tail to be at %d, got %d", test.tailPos, q.tail)
		}
		if cap(q.items) != test.expectedCap {
			t.Errorf("Expected cap of queue to be %. got %d", test.expectedCap, cap(q.items))
		}

	}
}
