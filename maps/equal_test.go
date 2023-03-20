package maps_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/generic/maps"
	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	assert := assert.New(t)

	type testCase struct {
		l, r map[int]string
		eq   bool
	}
	for _, tc := range []testCase{
		{
			l:  nil,
			r:  map[int]string{},
			eq: false,
		}, {
			l:  nil,
			r:  nil,
			eq: true,
		}, {
			l:  map[int]string{},
			r:  map[int]string{},
			eq: true,
		}, {
			l:  map[int]string{1: "foo"},
			r:  map[int]string{1: "foo"},
			eq: true,
		}, {
			l:  map[int]string{1: "foo", 2: "bar"},
			r:  map[int]string{1: "foo"},
			eq: false,
		},
	} {
		assert.Equalf(tc.eq, maps.Equal(tc.l, tc.r), "diff = %s", cmp.Diff(tc.l, tc.r))
	}
}
