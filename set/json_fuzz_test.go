package set_test

import (
	"encoding/binary"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/generic/set"
	"github.com/stretchr/testify/require"
)

var (
	minDate = time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli()
	maxDate = time.Date(9999, 12, 31, 12, 59, 59, 999000000, time.UTC).UnixMilli()
)

func FuzzJSON(f *testing.F) {
	f.Add([]byte("abcdefghijklmnopqrstuvwxyz1234567890-=]'"))
	f.Fuzz(func(t *testing.T, data []byte) {
		require := require.New(t)

		ints := set.NewOrdered[int64]()
		times := set.NewOrdered[time.Time]()

		for i := 0; i+8 < len(data); i = i + 8 {
			num := int64(binary.LittleEndian.Uint64(data[i : i+8]))
			ints.Add(num)
			if num < minDate {
				num = num % minDate
			} else if num > maxDate {
				num = num % maxDate
			}
			times.Add(time.UnixMilli(num))
		}

		// int
		{
			sl := ints.Values()

			marshalled, err := json.Marshal(ints)
			require.NoError(err)

			slMarshalled, err := json.Marshal(sl)
			require.NoError(err)

			if diff := cmp.Diff(string(slMarshalled), string(marshalled)); diff != "" {
				t.Fatalf(
					"not equal. diff = %s, slice = %s, OrderedSet = %s",
					diff, slMarshalled, marshalled,
				)
			}

			back := set.NewOrdered[int64]()
			require.NoError(json.Unmarshal(marshalled, back))
			backMarshalled, _ := json.Marshal(back)
			if diff := cmp.Diff(string(marshalled), string(backMarshalled)); diff != "" {
				t.Fatalf(
					"not equal. original = %s, slice = %s, marshalled and unmarshalled = %s",
					diff, marshalled, backMarshalled,
				)
			}

		}

		// time.Time
		{
			sl := times.Values()

			marshalled, err := json.Marshal(times)
			require.NoError(err)

			slMarshalled, err := json.Marshal(sl)
			require.NoError(err)
			if diff := cmp.Diff(string(slMarshalled), string(marshalled)); diff != "" {
				t.Fatalf(
					"not equal. diff = %s, slice = %s, OrderedSet = %s",
					diff, slMarshalled, marshalled,
				)
			}

			back := set.NewOrdered[time.Time]()
			require.NoError(json.Unmarshal(marshalled, back))
			backMarshalled, _ := json.Marshal(back)
			if diff := cmp.Diff(string(marshalled), string(backMarshalled)); diff != "" {
				t.Fatalf(
					"not equal. original = %s, slice = %s, marshalled and unmarshalled = %s",
					diff, marshalled, backMarshalled,
				)
			}
		}
	})

}
