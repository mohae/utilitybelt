// Queue implements a queue that will grow as needed, minimize unnecessary growth, and
// can be safely accessed from multiple routines without introducing race conditions.
// If necessary, the growth of the queue can be capped to a configurable size. When a
// maximum capacity for the queue is defined, an error will occur if the queue is full
// and an attempt is made to add another item to the queue.
//
// A queue is created with a minimum length and an optional maximum size (capacity).
// If the max size of the queue == 0, the queue will be unbounded. The growth rate of
// the queue is similar to that of a slice. When the queue grows, all items in the
// queue are shifted so that the head of the queue points to the first element in the
// queue.
//
// Before an item is enqueued, the queue is checked to see if thie new item
// will cause it to grow. If the tail == length, growth may occur. If the
// head of the queue is past a certain point in the queue, which is currently
// calculated using a percentage, the items in the queue will be shifted to start at
// the beginning of the slice, instead of growing the slice. The queue's head and tail
// will then be updated to reflect the shift.
//
// After dequeuing an item, the head position will be checked. If the queue
// io empty, head > tail, head and tail will be set to 0. This allows for
// efficient reuse of the queue without having to check to see if the queue items
// should be shifted or the queue should be grown.
//
// Once a queue grows, it will not be shrunk. This behavior may change in the future.
//
// All publicly exposed methods on the queue use locking to protect the queue from
// race conditions in situations where the queue is being accessed concurrently.
// Unexposed methods do not do any locking/unlocking since it is expected that the
// calling function has already obtained the lock and will release it as appropriate.
package queue

import (
	"fmt"
	"sync"
)

// shiftPercent is the default value for shifting the queue items to the
// front of the queue instead of growing the queue. If at least the % of
// the items have been removed from the queue, the items in the queue will
// be shifted to make room; otherwise the queue will grow
var shiftPercent = 20

// Queue represents a queue and everything needed to manage it. The preferred method
// for creating a new Queue is to use the New() func.
type Queue struct {
	mu           sync.Mutex
	items        []interface{}
	head         int // current item in queue
	tail         int // tail is the next insert point. last item is tail - 1
	length       int // the current length (cap) of the queue,
	maxCap       int // if > 0, the queue's cap cannot grow beyond this value
	shiftPercent int // the % of items that need to be removed before shifting occurs
}

// New returns an empty queue with a capacity equal to the recieved size value. If
// maxCap is > 0, the queue will not grow larger than maxCap; if it is at maxCap
// and growth is requred to enqueue an item, an error will occur.
func New(size, maxCap int) *Queue {
	return &Queue{items: make([]interface{}, size, size), length: size, maxCap: maxCap, shiftPercent: shiftPercent}
}

// Enqueue: adds an item to the queue. If adding the item requires growing
// the queue, the queue will either be shifted, to make room at the end of the queue
// or it will grow. If the queue cannot be grown, an error will be returned.
func (q *Queue) Enqueue(item interface{}) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	// See if it needs to grow
	if q.tail == q.length {
		shifted := q.shift()
		// if we weren't able to make room by shifting, grow the queue/
		if !shifted {
			err := q.grow()
			if err != nil {
				return err
			}
		}
	}
	q.items[q.tail] = item
	q.tail++
	return nil
}

// Dequeue removes an item from the queue. If the removal of the item empties
// the queue, the head and tail will be set to 0.
func (q *Queue) Dequeue() interface{} {
	q.mu.Lock()
	i := q.items[q.head]
	q.head++
	if q.head > q.tail {
		q.mu.Unlock()
		q.Reset()
		return i
	}
	q.mu.Unlock()
	return i
}

// IsEmpty returns whether or not the queue is empty
func (q *Queue) IsEmpty() bool {
	q.mu.Lock()
	if q.tail == 0 || q.head == q.tail {
		q.mu.Unlock()
		return true
	}
	q.mu.Unlock()
	return false
}

func (q *Queue) Tail() int {
	return q.tail
}

func (q *Queue) Head() int {
	return q.head
}

// shift: if shiftPercent items have been removed from the queue, the remaining items
// in the queue will be shifted to element 0-n, where n is the number of remaining
// items in the queue. Returns whether or not a shift occurred
func (q *Queue) shift() bool {
	if q.head <= (q.length*q.shiftPercent)/100 {
		return false
	}
	copy(q.items, q.items[q.head:q.tail])
	// set the pointers to the correct position
	q.tail = q.tail - q.head
	q.head = 0
	return true
}

// grow grows the slice using an algorithm similar to growSlice(). This is a bit slower
// than relying on slice's automatic growth, but allows for capacity enforcement w/o
// growing the slice cap beyond the configured maxCap, if applicable.
//
// Since a temporary slice is created to store the current queue, all items in queue
// are automatically shifted
func (q *Queue) grow() error {
	if q.length == q.maxCap && q.maxCap > 0 {
		return fmt.Errorf("groweQueue: cannot grow beyond max capacity of %d", q.maxCap)
	}
	if q.length < 1024 {
		q.length += q.length
	} else {
		q.length += q.length / 4
	}
	// If the maxCap is set, cannot grow it beyond that
	if q.length > q.maxCap && q.maxCap > 0 {
		q.length = q.maxCap
	}
	// grow the slice
	l := q.tail - q.head
	tmp := make([]interface{}, l, l)
	copy(tmp, q.items[q.head:q.tail])
	q.items = make([]interface{}, q.length, q.length)
	copy(q.items, tmp)
	q.tail = l
	q.head = 0
	return nil
}

// Reset resets the queue; head and tail point to element 0.
func (q *Queue) Reset() {
	q.mu.Lock()
	q.head = 0
	q.tail = 0
	q.mu.Unlock()
}
