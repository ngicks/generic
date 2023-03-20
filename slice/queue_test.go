package slice_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/generic/slice"
)

func TestQueue(t *testing.T) {
	testQueue(t, slice.Queue[int]{})
	testQueue(t, nil)
}

func testQueue(t *testing.T, queue slice.Queue[int]) {
	if queue.Len() != 0 {
		t.Fatalf("wrong len: expected to be %d but is %d", 0, queue.Len())
	}

	if v, popped := queue.Pop(); popped {
		t.Fatalf("incorrect behavior: popped is expected to be false but [%t] with [%d]", popped, v)
	}

	queue.Push(6)
	queue.Push(3)
	queue.Push(1)
	queue.Push(2)
	queue.Push(4)
	queue.Push(5)

	if queue.Len() != 6 {
		t.Fatalf("wrong len: expected to be %d but is %d", 6, queue.Len())
	}
	expected := []int{6, 3, 1, 2, 4, 5}
	if cmp.Diff(([]int)(queue), expected) != "" {
		t.Fatalf("incorrect push behavior: expected to be %v, but is %v", expected, queue)
	}

	popped := []int{}
	var v int
	v, _ = queue.Pop()
	popped = append(popped, v)
	v, _ = queue.Pop()
	popped = append(popped, v)
	v, _ = queue.Pop()
	popped = append(popped, v)
	v, _ = queue.Pop()
	popped = append(popped, v)
	v, _ = queue.Pop()
	popped = append(popped, v)
	v, _ = queue.Pop()
	popped = append(popped, v)

	if cmp.Diff(popped, expected) != "" {
		t.Fatalf("incorrect pop behavior: expected to be %v, but is %v", expected, popped)
	}
	if queue.Len() != 0 {
		t.Fatalf("wrong len: expected to be %d but is %d", 0, queue.Len())
	}

	queue.Push(6)
	queue.Push(3)
	queue.Push(1)
	queue.Push(2)
	queue.Push(4)
	queue.Push(5)
	cloned := queue.Clone()
	if cmp.Diff(queue, cloned) != "" {
		t.Fatalf("incorrect clone behavior: expected to be %v, but is %v", queue, cloned)
	}
	cloned[1] = 13
	if cloned[1] == queue[1] {
		t.Fatalf("incorrect clone behavior: expected to be %v, but is %v", cloned[1], queue[1])
	}

	queue = []int{}

	queue.Prepend([]int{1, 2, 3, 4, 5}...)
	if expected := []int{1, 2, 3, 4, 5}; cmp.Diff(expected, ([]int)(queue)) != "" {
		t.Fatalf("incorrect prepend behavior: expected to be %v, but is %v", expected, queue)
	}

	queue.Prepend([]int{11, 12, 13, 14, 15}...)
	if expected := []int{11, 12, 13, 14, 15, 1, 2, 3, 4, 5}; cmp.Diff(expected, ([]int)(queue)) != "" {
		t.Fatalf("incorrect prepend behavior: expected to be %v, but is %v", expected, queue)
	}

}
