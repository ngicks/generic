package set

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/generic/slice"
)

var _ SetLike[int] = &Set[int]{}
var _ SetLike[int] = &OrderedSet[int]{}

var utc, jst *time.Location

func init() {
	var err error
	utc, err = time.LoadLocation("")
	if err != nil {
		panic(err)
	}
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
}

func TestSet(t *testing.T) {
	set := New[time.Time]()
	testSetWithTimeKey(t, set)
}

func TestOrderedSet(t *testing.T) {
	set := NewOrdered[time.Time]()
	testSetWithTimeKey(t, set)
}

func TestOrderedSetOrdering(t *testing.T) {
	set := NewOrdered[int]()

	set.Add(5)
	set.Add(4)
	set.Add(3)
	set.Add(2)
	set.Add(1)

	expected := []int{5, 4, 3, 2, 1}

	assertInsertionOrder := func(t *testing.T, set SetLike[int], expected []int) {
		collected := set.Values()
		if cmp.Diff(expected, collected) != "" {
			t.Fatalf("must be deeply equal.\nexpected = %+v\nactual = %+v\n", expected, collected)
		}

		collected = collected[:0]

		set.ForEach(func(v, _ int) {
			collected = append(collected, v)
		})
		if cmp.Diff(expected, collected) != "" {
			t.Fatalf("must be deeply equal.\nexpected = %+v\nactual = %+v\n", expected, collected)
		}
	}

	for i := 0; i < 100; i++ {
		assertInsertionOrder(t, set, expected)
	}
}

func testSetWithTimeKey(t *testing.T, set SetLike[time.Time]) {
	now := time.Now()
	year, month, day := now.Date()
	hour, minute, second := now.Clock()
	nanosecond := now.Nanosecond()

	tgt := time.Date(year, month, day, hour, minute, second, nanosecond, utc)
	tgtInOtherLocation := time.Date(year, month, day, hour, minute, second, nanosecond, jst)

	if set.Len() != 0 {
		t.Fatalf("wrong len")
	}
	set.Add(tgt)
	if set.Len() != 1 {
		t.Fatalf("wrong len")
	}
	set.Add(tgtInOtherLocation)
	if set.Len() != 2 {
		t.Fatalf("wrong len")
	}
	set.Add(tgt)
	if set.Len() != 2 {
		t.Fatalf("wrong len")
	}
	if !set.Has(tgt) {
		t.Fatalf("must have")
	}
	if set.Has(now) {
		t.Fatalf("must not have")
	}

	d := set.Values()

	if !(slice.Has(d, tgt) && slice.Has(d, tgtInOtherLocation)) {
		t.Fatalf("must contain")
	}

	d = d[:0]
	idxSl := make([]int, 0)
	set.ForEach(func(v time.Time, idx int) {
		d = append(d, v)
		idxSl = append(idxSl, idx)
	})
	if !isSerialSlice(idxSl, 1) || idxSl[0] != 0 {
		t.Fatalf("index must be 0,1,2,3... but = %+v", idxSl)
	}
	if !(slice.Has(d, tgt) && slice.Has(d, tgtInOtherLocation)) {
		t.Fatalf("must contain")
	}

	if set.Delete(now); set.Len() != 2 {
		t.Fatalf("wrong len")
	}
	if set.Delete(tgt); set.Len() != 1 {
		t.Fatalf("wrong len")
	}

	set.Add(now.Add(time.Second))
	set.Add(now.Add(2 * time.Second))
	set.Add(now.Add(3 * time.Second))
	set.Clear()
	if set.Len() != 0 {
		t.Fatalf("wrong len")
	}
}

func isSerialSlice(sl []int, width int) bool {
	prev := sl[0]
	for _, v := range sl[1:] {
		if v-prev != width {
			return false
		}
	}
	return true
}
