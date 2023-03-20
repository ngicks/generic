package slice_test

import (
	"runtime"
	"sync/atomic"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/generic/slice"
)

func TestDequeue(t *testing.T) {
	testDequeue(t, make(slice.Deque[int], 0))
	testDequeue(t, slice.Deque[int]{})
	testDequeue(t, nil)
}

func testDequeue(t *testing.T, deque slice.Deque[int]) {
	if deque.Len() != 0 {
		t.Fatalf("wrong len: expected to be %d but is %d", 0, deque.Len())
	}

	if v, popped := deque.Pop(); popped {
		t.Fatalf("incorrect behavior: popped is expected to be false but [%t] with [%d]", popped, v)
	}
	if v, popped := deque.PopBack(); popped {
		t.Fatalf("incorrect behavior: popped is expected to be false but [%t] with [%d]", popped, v)
	}
	if v, popped := deque.PopFront(); popped {
		t.Fatalf("incorrect behavior: popped is expected to be false but [%t] with [%d]", popped, v)
	}

	deque.PushFront(1)
	deque.PushBack(2)
	deque.PushFront(3)
	deque.PushBack(4)
	deque.Push(5)
	deque.PushFront(6)

	if deque.Len() != 6 {
		t.Fatalf("wrong len: expected to be %d but is %d", 6, deque.Len())
	}

	{
		expected := []int{6, 3, 1, 2, 4, 5}
		if cmp.Diff(([]int)(deque), expected) != "" {
			t.Fatalf("incorrect push behavior: expected to be %v, but is %v", expected, deque)
		}
	}

	popped := []int{}
	var v int
	v, _ = deque.PopFront()
	popped = append(popped, v)
	v, _ = deque.Pop()
	popped = append(popped, v)
	v, _ = deque.PopBack()
	popped = append(popped, v)
	v, _ = deque.PopFront()
	popped = append(popped, v)
	v, _ = deque.PopBack()
	popped = append(popped, v)
	v, _ = deque.PopFront()
	popped = append(popped, v)

	if deque.Len() != 0 {
		t.Fatalf("wrong len: expected to be %d but is %d", 0, deque.Len())
	}
	{
		expected := []int{6, 5, 4, 3, 2, 1}
		if cmp.Diff(popped, expected) != "" {
			t.Fatalf("incorrect pop behavior: expected to be %v, but is %v", expected, popped)
		}
	}

	deque.Append([]int{1, 2, 3, 4, 5}...)
	if expected := []int{1, 2, 3, 4, 5}; cmp.Diff(expected, []int(deque)) != "" {
		t.Fatalf("incorrect append behavior: expected to be %v, but is %v", expected, deque)
	}
	deque.Append([]int{9, 10, 11, 12}...)
	if expected := []int{1, 2, 3, 4, 5, 9, 10, 11, 12}; cmp.Diff(expected, []int(deque)) != "" {
		t.Fatalf("incorrect append behavior: expected to be %v, but is %v", expected, deque)
	}

	deque = []int{}

	deque.Prepend([]int{1, 2, 3, 4, 5}...)
	if expected := []int{1, 2, 3, 4, 5}; cmp.Diff(expected, ([]int)(deque)) != "" {
		t.Fatalf("incorrect prepend behavior: expected to be %v, but is %v", expected, deque)
	}

	deque.Prepend([]int{11, 12, 13, 14, 15}...)
	if expected := []int{11, 12, 13, 14, 15, 1, 2, 3, 4, 5}; cmp.Diff(expected, ([]int)(deque)) != "" {
		t.Fatalf("incorrect prepend behavior: expected to be %v, but is %v", expected, deque)
	}

	cloned := deque.Clone()
	if cmp.Diff(deque, cloned) != "" {
		t.Fatalf("incorrect clone behavior: expected to be %v, but is %v", deque, cloned)
	}
	cloned[1] = 13
	if cloned[1] == deque[1] {
		t.Fatalf("incorrect clone behavior: expected to be %v, but is %v", cloned[1], deque[1])
	}

	if g, ok := deque.Get(1); !ok || g != deque[1] {
		t.Fatalf("incorrect get behavior: expected to be %v, but is %v", deque[1], g)
	}
	if g, ok := deque.Get(100); ok || g != 0 {
		t.Fatalf("incorrect get behavior: expected to be %v, but is %v", deque[1], g)
	}

	deque = []int{}
	deque.Append([]int{0, 1, 2, 3, 4}...)
	deque.Insert(0, 150)
	if expected := []int{150, 0, 1, 2, 3, 4}; cmp.Diff(expected, []int(deque)) != "" {
		t.Fatalf("incorrect append behavior: expected to be %v, but is %v", expected, deque)
	}
	deque.Insert(4, 200)
	if expected := []int{150, 0, 1, 2, 200, 3, 4}; cmp.Diff(expected, []int(deque)) != "" {
		t.Fatalf("incorrect append behavior: expected to be %v, but is %v", expected, deque)
	}

	func() {
		defer func() {
			recv := recover()
			if recv == nil {
				t.Fatalf("must panic")
			}
		}()
		deque.Insert(uint(deque.Len()+1), 120)
	}()
}

func TestDeque_popped_element_can_be_GCed(t *testing.T) {
	// Do not make this like slice.Deque[*int]{}
	// I dunno much about it but seemingly
	// Go compiler optimizes int to be like uint8.
	// And it causes this test to be flaky and unable to pass.
	d := slice.Deque[*string]{}

	// called is count for SetFinalizer ever called.
	var called int64

	for _, v := range []string{"foo", "bar", "baz", "qux", "quux"} {
		copied := v
		runtime.SetFinalizer(&copied, func(*string) {
			atomic.AddInt64(&called, 1)
		})
		d.Push(&copied)
	}

	if loaded := atomic.LoadInt64(&called); loaded != 0 {
		t.Fatalf("finalizer must not be called at this moment, is %d", loaded)
	}

	d.PopBack()
	runtime.GC()
	for _, v := range d {
		// keeping alive references.
		runtime.KeepAlive(v)
	}
	runtime.GC()
	if loaded := atomic.LoadInt64(&called); loaded != 1 {
		t.Fatalf("finalizer must be called once, but is %d", loaded)
	}

	d.PopBack()
	runtime.GC()
	for _, v := range d {
		runtime.KeepAlive(v)
	}
	runtime.GC()
	if loaded := atomic.LoadInt64(&called); loaded != 2 {
		t.Fatalf("finalizer must be called twice, but is %d", loaded)
	}

	d.PopFront()
	runtime.GC()
	for _, v := range d {
		runtime.KeepAlive(v)
	}
	runtime.GC()
	if loaded := atomic.LoadInt64(&called); loaded != 3 {
		t.Fatalf("finalizer must be called 3 times, but is %d", loaded)
	}

	for _, v := range d {
		runtime.KeepAlive(v)
	}
}
