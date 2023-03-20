package set

import (
	"bytes"
	"container/list"
	"encoding/json"
	"fmt"

	"github.com/buger/jsonparser"
	"github.com/mailru/easyjson/jwriter"
)

// OrderedSet is same as Set but remembers insertion order.
type OrderedSet[T comparable] struct {
	order  *list.List
	eleMap map[T]*list.Element
}

func NewOrdered[T comparable]() *OrderedSet[T] {
	return &OrderedSet[T]{
		order:  list.New(),
		eleMap: make(map[T]*list.Element),
	}
}

func (s *OrderedSet[T]) Len() int {
	return len(s.eleMap)
}

func (s *OrderedSet[T]) Add(v T) {
	_, has := s.eleMap[v]
	if !has {
		ele := s.order.PushBack(v)
		s.eleMap[v] = ele
	}
}

func (s *OrderedSet[T]) Clear() {
	s.order = s.order.Init()
	s.eleMap = make(map[T]*list.Element)
}

func (s *OrderedSet[T]) Delete(v T) (deleted bool) {
	ele, deleted := s.eleMap[v]
	if deleted {
		s.order.Remove(ele)
		delete(s.eleMap, v)
	}
	return
}

// ForEach iterates over set and invoke f with each elements.
// The order is FIFO. Add with an existing v does not update the order.
func (s *OrderedSet[T]) ForEach(f func(v T, idx int)) {
	var idx int
	for next := s.order.Front(); next != nil; next = next.Next() {
		f(next.Value.(T), idx)
		idx++
	}
}

func (s *OrderedSet[T]) Has(v T) (has bool) {
	_, has = s.eleMap[v]
	return
}

func (s *OrderedSet[T]) Values() []T {
	sl := make([]T, 0)
	s.ForEach(func(v T, _ int) {
		sl = append(sl, v)
	})
	return sl
}

func (s OrderedSet[T]) MarshalJSON() ([]byte, error) {
	var writer jwriter.Writer

	writer.RawByte('[')

	firstElement := true
	for next := s.order.Front(); next != nil; next = next.Next() {
		if !firstElement {
			writer.RawByte(',')
		}
		writer.Raw(json.Marshal(next.Value.(T)))
		if writer.Error != nil {
			return nil, writer.Error
		}
		firstElement = false
	}

	writer.RawByte(']')

	buf := new(bytes.Buffer)
	buf.Grow(writer.Size())
	if _, err := writer.DumpTo(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *OrderedSet[T]) UnmarshalJSON(data []byte) error {
	if data[0] != '[' {
		return fmt.Errorf(
			"OrderedSet: can not marshal values other than array into OrderedSet. input = %s",
			data,
		)
	}

	s.Clear()

	var (
		backed       byte
		savedOffset  int
		unmarshalErr error
	)
	_, err := jsonparser.ArrayEach(
		data,
		func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			if dataType == jsonparser.String {
				// get back double-quotations
				value = data[offset-2 : offset+len(value)]
			}

			var val T
			unmarshalErr = json.Unmarshal(value, &val)
			if unmarshalErr != nil {
				// I do not know why, but offset is (starting point of value + 1).
				delimIdx := offset + len(value)
				if dataType == jsonparser.String {
					// We've extended value length by 2 for string type.
					delimIdx = delimIdx - 2
				}
				// stop iteration.
				// jsonparser does not provide a way to stop its iteration on an array.
				// see https://github.com/buger/jsonparser/issues/255.
				backed = data[delimIdx]
				data[delimIdx] = ']'
				savedOffset = delimIdx

				return
			}
			s.Add(val)
		},
	)

	if unmarshalErr != nil {
		data[savedOffset] = backed
		s.Clear()
		return unmarshalErr
	}
	return err
}
