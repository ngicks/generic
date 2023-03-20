package maps_test

import (
	"encoding/binary"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/generic/maps"
	"github.com/stretchr/testify/require"
)

func FuzzMaps(f *testing.F) {
	f.Add([]byte("abcdefghijklmnopqrstuvwxyz1234567890-=]dih3y9ra';,3q.dopjzp0ju-wqk[p]'"))
	f.Fuzz(func(t *testing.T, data []byte) {
		require := require.New(t)

		org := map[int64]int64{}

		for i := 0; i+8 < len(data); i = i + 8 {
			num1 := int64(binary.LittleEndian.Uint32(data[i : i+4]))
			num2 := int64(binary.LittleEndian.Uint32(data[i+4 : i+8]))
			org[num1] = num2
		}

		cloned := maps.Clone(org)
		require.True(maps.Equal(org, cloned))

		var someKey int64
		for k := range org {
			someKey = k
			break
		}

		cloned[someKey] = -1
		require.False(maps.Equal(org, cloned))
	})
}

func TestCloneFiltered(t *testing.T) {
	input := map[string]string{
		"foo": "foo",
		"bar": "bar",
		"baz": "baz",
		"qux": "qux",
	}

	filter := func(k, v string) bool {
		return strings.HasPrefix(v, "b")
	}

	if diff := cmp.Diff(
		map[string]string{"bar": "bar", "baz": "baz"},
		maps.CloneSelected(input, filter),
	); diff != "" {
		t.Fatalf("diff = %s", diff)
	}

	if diff := cmp.Diff(
		map[string]string{"foo": "foo", "qux": "qux"},
		maps.CloneExcluded(input, filter),
	); diff != "" {
		t.Fatalf("diff = %s", diff)
	}

}
