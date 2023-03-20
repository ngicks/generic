package set_test

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ngicks/generic/set"
	"github.com/stretchr/testify/require"
)

type sample struct {
	A int
	B float64
	C string
	D bool
}

func TestJSON(t *testing.T) {
	require := require.New(t)

	structSet := set.NewOrdered[sample]()
	bin := []byte(`[{"A":123,"B":0.25,"C":"foo","D":true},{"A":224,"B":2317.09342,"C":"bar","D":false}]`)
	require.NoError(json.Unmarshal(bin, structSet))
	back, err := json.Marshal(structSet)
	require.NoError(err)
	if diff := cmp.Diff(string(bin), string(back)); diff != "" {
		t.Fatalf(
			"encoded back to different string. diff = %s, input = %s, decoded and encoded back = %s",
			diff, bin, back,
		)
	}
}

func TestJSON_error(t *testing.T) {
	require := require.New(t)
	intSet := set.NewOrdered[int]()

	bin := []byte(`[123, "bar", "baz"]`)
	require.Error(json.Unmarshal(bin, &intSet))
	// making sure unmarshal does not mutate input byte slice.
	if diff := cmp.Diff(`[123, "bar", "baz"]`, string(bin)); diff != "" {
		t.Fatalf("passed data is mutated incorrectly. diff = %s", diff)
	}

	strSet := set.NewOrdered[string]()
	bin = []byte(`["123",456,789]`)
	require.Error(json.Unmarshal(bin, strSet))
	if diff := cmp.Diff(`["123",456,789]`, string(bin)); diff != "" {
		t.Fatalf("passed data is mutated incorrectly. diff = %s", diff)
	}

	structSet := set.NewOrdered[sample]()
	bin = []byte(`[{"A":123,"B":0.25,"C":"foo","D":true},{"A":224,"B":2317.09342,"C":"bar","D":false},"foo",{"A":224,"B":2317.09342,"C":"bar","D":false}]`)
	require.Error(json.Unmarshal(bin, structSet))
	if diff := cmp.Diff(
		`[{"A":123,"B":0.25,"C":"foo","D":true},{"A":224,"B":2317.09342,"C":"bar","D":false},"foo",{"A":224,"B":2317.09342,"C":"bar","D":false}]`,
		string(bin),
	); diff != "" {
		t.Fatalf("passed data is mutated incorrectly. diff = %s", diff)
	}
}
