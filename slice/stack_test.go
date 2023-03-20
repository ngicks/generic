package slice_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/generic/slice"
)

func TestStack(t *testing.T) {
	testStack(t, slice.Stack[int]{})
	testStack(t, nil)
}

func testStack(t *testing.T, stack slice.Stack[int]) {
	if stack.Len() != 0 {
		t.Fatalf("wrong len: expected to be %d but is %d", 0, stack.Len())
	}
	if v, popped := stack.Pop(); popped {
		t.Fatalf("incorrect behavior: popped is expected to be false but [%t] with [%d]", popped, v)
	}

	stack.Push(6)
	stack.Push(3)
	stack.Push(1)
	stack.Push(2)
	stack.Push(4)
	stack.Push(5)

	if stack.Len() != 6 {
		t.Fatalf("wrong len: expected to be %d but is %d", 6, stack.Len())
	}
	{
		expected := []int{6, 3, 1, 2, 4, 5}
		if cmp.Diff(([]int)(stack), expected) != "" {
			t.Fatalf("incorrect push behavior: expected to be %v, but is %v", expected, stack)
		}
	}

	popped := []int{}
	var v int
	v, _ = stack.Pop()
	popped = append(popped, v)
	v, _ = stack.Pop()
	popped = append(popped, v)
	v, _ = stack.Pop()
	popped = append(popped, v)
	v, _ = stack.Pop()
	popped = append(popped, v)
	v, _ = stack.Pop()
	popped = append(popped, v)
	v, _ = stack.Pop()
	popped = append(popped, v)

	if stack.Len() != 0 {
		t.Fatalf("wrong len: expected to be %d but is %d", 0, stack.Len())
	}
	{
		reversed := []int{6, 3, 1, 2, 4, 5}
		expected := make([]int, len(reversed))
		for idx, v := range reversed {
			expected[len(reversed)-(idx+1)] = v
		}
		if cmp.Diff(popped, expected) != "" {
			t.Fatalf("incorrect pop behavior: expected to be %v, but is %v", expected, popped)
		}
	}

	stack.Append([]int{1, 2, 3, 4, 5}...)
	if expected := []int{1, 2, 3, 4, 5}; cmp.Diff(expected, []int(stack)) != "" {
		t.Fatalf("incorrect append behavior: expected to be %v, but is %v", expected, stack)
	}
	stack.Append([]int{9, 10, 11, 12}...)
	if expected := []int{1, 2, 3, 4, 5, 9, 10, 11, 12}; cmp.Diff(expected, []int(stack)) != "" {
		t.Fatalf("incorrect append behavior: expected to be %v, but is %v", expected, stack)
	}

	cloned := stack.Clone()
	if cmp.Diff(stack, cloned) != "" {
		t.Fatalf("incorrect clone behavior: expected to be %v, but is %v", stack, cloned)
	}
	cloned[1] = 13
	if cloned[1] == stack[1] {
		t.Fatalf("incorrect clone behavior: expected to be %v, but is %v", cloned[1], stack[1])
	}
}
