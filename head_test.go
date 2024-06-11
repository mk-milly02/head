package head_test

import (
	"head"
	"reflect"
	"testing"
)

func TestReadFirstNLines(t *testing.T) {
	t.Parallel()

	subtests := []struct {
		text []byte
		n    int
		want []string
	}{
		{
			text: []byte("ping mpls pseudowire/r/nping vpls vsi"),
			n:    0,
			want: []string{"ping mpls pseudowire", "ping vpls vsi"},
		},
		{
			text: []byte("ping mpls pseudowire/r/nping vpls vsi"),
			n:    1,
			want: []string{"ping mpls pseudowire"},
		},
	}

	for _, st := range subtests {
		if got := head.ReadFirstNLines(st.text, st.n); reflect.DeepEqual(got, st.want) {
			t.Errorf("wanted %v : got %v", st.want, got)
		}
	}
}
