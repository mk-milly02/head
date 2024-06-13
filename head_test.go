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
			text: []byte("ping mpls pseudowire\nping vpls vsi"),
			n:    0,
			want: nil,
		},
		{
			text: []byte("ping mpls pseudowire\nping vpls vsi"),
			n:    1,
			want: []string{"ping mpls pseudowire\n"},
		},
	}

	for _, st := range subtests {
		if got := head.ReadFirstNLines(st.text, st.n); !reflect.DeepEqual(got, st.want) {
			t.Errorf("wanted %v : got %v", st.want, got)
		}
	}
}

/*func TestCLI(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("./head")
	testCases := []struct {
		command string
		args    []string
		want    string
	}{
		{
			command: "go run cmd/main.go test.txt",
			args:    []string{"./test.txt"},
			want: `The Project Gutenberg eBook of The Art of War

				This ebook is for the use of anyone anywhere in the United States and
				most other parts of the world at no cost and with almost no restrictions
				whatsoever. You may copy it, give it away or re-use it under the terms
				of the Project Gutenberg License included with this ebook or online
				at www.gutenberg.org. If you are not located in the United States,
				you will have to check the laws of the country where you are located
				before using this eBook.`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.command, func(t *testing.T) {
			cmd.Args = tC.args
			got, err := cmd.CombinedOutput()
			if err != nil {
				t.Error(err)
			} else {
				if tC.want != string(got) {
					t.Errorf("wanted: \n%s \n\ngot: \n%s", tC.want, string(got))
				}

			}
		})
	}
}*/

func TestReadFirstCBytes(t *testing.T) {
	type args struct {
		text []byte
		c    int
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name: "hydrogen",
			args: args{
				text: []byte("Hello"),
				c:    2,
			},
			wantResult: "He",
		},
		{
			name: "helium",
			args: args{
				text: []byte("Hello, world!"),
				c:    10,
			},
			wantResult: "Hello, wor",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := head.ReadFirstCBytes(tt.args.text, tt.args.c); gotResult != tt.wantResult {
				t.Errorf("ReadFirstCBytes() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
