package slice_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/generic/slice"
)

func TestFunc(t *testing.T) {
	sl := []int{1, 2, 3, 4, 1, 5}

	if !slice.Has(sl, 1) {
		t.Fatalf("must be true")
	}
	if slice.Has(sl, 10) {
		t.Fatalf("must be false")
	}
	if slice.Has(nil, 10) {
		t.Fatalf("returned value must be false if input slice is nil")
	}

	if idx, ok := slice.Find(sl, func(v int) bool { return v > 2 }); !ok || idx != 3 {
		t.Fatalf("index must be %d, but %d", 3, idx)
	}
	if idx, ok := slice.Find(sl, func(v int) bool { return v == 6 }); ok {
		t.Fatalf("returned ok must be false if no value matches to given predicate, but found %d", idx)
	}
	if _, ok := slice.Find(nil, func(v int) bool { return v == 1 }); ok {
		t.Fatalf("returned ok must be false if any of input is nil")
	}
	if _, ok := slice.Find(sl, nil); ok {
		t.Fatalf("returned ok must be false if any of input is nil")
	}
	if _, ok := slice.Find[int](nil, nil); ok {
		t.Fatalf("returned ok must be false if any of input is nil")
	}

	if idx, ok := slice.FindLast(sl, func(v int) bool { return v > 2 }); !ok || idx != 5 {
		t.Fatalf("index must be %d, but %d", 5, idx)
	}
	if idx, ok := slice.FindLast(sl, func(v int) bool { return v == 6 }); ok {
		t.Fatalf("returned ok must be false if no value matches to given predicate, but found %d", idx)
	}
	if _, ok := slice.FindLast(nil, func(v int) bool { return v == 1 }); ok {
		t.Fatalf("returned ok must be false if any of input is nil")
	}
	if _, ok := slice.FindLast(sl, nil); ok {
		t.Fatalf("returned ok must be false if any of input is nil")
	}
	if _, ok := slice.FindLast[int](nil, nil); ok {
		t.Fatalf("returned ok must be false if any of input is nil")
	}

	if p := slice.Position(sl, func(v int) bool { return v == 1 }); p != 0 {
		t.Fatalf("must be 0, but %d", p)
	}
	if p := slice.Position(sl, func(v int) bool { return v == 6 }); p != -1 {
		t.Fatalf("must be -1, but %d", p)
	}
	if p := slice.Position(nil, func(v int) bool { return v == 1 }); p != -1 {
		t.Fatalf("must be -1, but %d", p)
	}
	if p := slice.Position(sl, nil); p != -1 {
		t.Fatalf("must be -1, but %d", p)
	}
	if p := slice.Position[int](nil, nil); p != -1 {
		t.Fatalf("must be -1, but %d", p)
	}

	if p := slice.PositionLast(sl, func(v int) bool { return v == 1 }); p != 4 {
		t.Fatalf("must be 4, but %d", p)
	}
	if p := slice.PositionLast(sl, func(v int) bool { return v == 6 }); p != -1 {
		t.Fatalf("must be -1, but %d", p)
	}
	if p := slice.PositionLast(nil, func(v int) bool { return v == 1 }); p != -1 {
		t.Fatalf("must be -1, but %d", p)
	}
	if p := slice.PositionLast(sl, nil); p != -1 {
		t.Fatalf("must be -1, but %d", p)
	}
	if p := slice.PositionLast[int](nil, nil); p != -1 {
		t.Fatalf("must be -1, but %d", p)
	}
}

func TestInsert(t *testing.T) {
	sl := []int{1, 2, 3}
	sl = slice.Insert(sl, 1, 4)
	if diff := cmp.Diff([]int{1, 4, 2, 3}, sl); diff != "" {
		t.Fatalf("not equal. expected = %+v, diff = %s", []int{1, 4, 2, 3}, diff)
	}
	sl = slice.Insert(sl, 4, 5)
	if cmp.Diff([]int{1, 4, 2, 3, 5}, sl) != "" {
		t.Fatalf("not equal. expected = %+v, actual = %+v", []int{1, 4, 2, 3, 5}, sl)
	}

	sl = []int{1, 2, 3}

	sl = slice.Insert(sl, 3, 25)
	if cmp.Diff([]int{1, 2, 3, 25}, sl) != "" {
		t.Fatalf("not equal. expected = %+v, actual = %+v", []int{1, 2, 3, 25}, sl)
	}
	sl = slice.Insert(sl, 0, 50)
	if cmp.Diff([]int{50, 1, 2, 3, 25}, sl) != "" {
		t.Fatalf("not equal. expected = %+v, actual = %+v", []int{50, 1, 2, 3, 25}, sl)
	}

	func() {
		defer func() {
			recv := recover()
			if recv == nil {
				t.Fatalf("must panic")
			}
		}()
		slice.Insert(sl, uint(len(sl)+1), 120)
	}()

	slAnother := slice.Insert(sl, 1, 1024)
	if cmp.Diff([]int{50, 1024, 1, 2, 3, 25}, slAnother) != "" {
		t.Fatalf("not equal. expected = %+v, actual = %+v", []int{50, 1024, 1, 2, 3, 25}, sl)
	}
}

func TestGet(t *testing.T) {
	sl := []int{1, 2, 3}

	for i := 0; i < 50; i++ {
		g, ok := slice.Get(sl, uint(i))
		if i < 3 {
			if g != i+1 || !ok {
				t.Fatalf("failed: %d", g)
			}
		} else {
			if g != 0 || ok {
				t.Fatalf("failed: %d", g)
			}
		}
	}
}

func TestPrepend(t *testing.T) {
	sl := []int{1, 2, 3}

	sl = slice.Prepend(sl, []int{5, 6, 7}...)

	if expected := []int{5, 6, 7, 1, 2, 3}; cmp.Diff(expected, sl) != "" {
		t.Fatalf("not equal. expected = %+v , actual = %+v", expected, sl)
	}

	sl = slice.Prepend(sl)
	if expected := []int{5, 6, 7, 1, 2, 3}; cmp.Diff(expected, sl) != "" {
		t.Fatalf("not equal. expected = %+v , actual = %+v", expected, sl)
	}
}

func TestClone(t *testing.T) {
	sl := []int{1, 2, 3}

	cloned := slice.Clone(sl)

	if cmp.Diff(sl, cloned) != "" {
		t.Fatalf("not equal. expected = %+v , actual = %+v", sl, cloned)
	}

	cloned[1] = 2000
	if sl[1] == cloned[1] {
		t.Fatalf(
			"cloned and original slices share a same underlying array."+
				" expected = %+v , actual = %+v",
			2, cloned[1],
		)
	}
}
