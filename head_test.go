package head_test

import (
	"head"
	"os/exec"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
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
		if got := head.ReadFirstNLines(st.text, st.n); !cmp.Equal(got, st.want) {
			t.Errorf("wanted %v : got %v", st.want, got)
		}
	}
}

func TestReadFirstCBytes(t *testing.T) {
	t.Parallel()

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
			if gotResult := head.ReadFirstCBytes(tt.args.text, tt.args.c); !cmp.Equal(tt.wantResult, gotResult) {
				t.Errorf("ReadFirstCBytes() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestMainStep1_1(t *testing.T) {
	t.Parallel()

	want := "The Project Gutenberg eBook of The Art of War\r\n    \r\nThis ebook is for the use of anyone anywhere in the United States and\r\nmost other parts of the world at no cost and with almost no restrictions\r\nwhatsoever. You may copy it, give it away or re-use it under the terms\r\nof the Project Gutenberg License included with this ebook or online\r\nat www.gutenberg.org. If you are not located in the United States,\r\nyou will have to check the laws of the country where you are located\r\nbefore using this eBook.\r\n\r\n"
	cmd := exec.Command("./cchead", "test.txt")

	got, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	} else {
		if !cmp.Equal(want, string(got)) {
			t.Errorf("\nwant: \n%s \n\ngot: \n%s", want, string(got))
		}

	}
}

func TestMainStep1_2(t *testing.T) {
	t.Parallel()

	want := "a\na\na\na\na\na\na\na\na\na\n"
	cmd := exec.Command("./cchead")
	cmd.Stdin = strings.NewReader("a")

	got, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	} else {
		if !cmp.Equal(want, string(got)) {
			t.Errorf("\nwant: \n%s \n\ngot: \n%s", want, string(got))
		}

	}
}

func TestMainStep2_1(t *testing.T) {
	t.Parallel()

	want := "The Project Gutenberg eBook of The Art of War\r\n"
	cmd := exec.Command("./cchead", "-n", "1", "test.txt")

	got, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	} else {
		if !cmp.Equal(want, string(got)) {
			t.Errorf("\nwant: \n%s \n\ngot: \n%s", want, string(got))
		}

	}
}

func TestMainStep2_2(t *testing.T) {
	t.Parallel()

	want := "Hello, World\n"
	cmd := exec.Command("./cchead", "-n", "3", "test2.txt")

	got, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	} else {
		if !cmp.Equal(want, string(got)) {
			t.Errorf("\nwant: \n%s \n\ngot: \n%s", want, string(got))
		}

	}
}

func TestMainStep3_1(t *testing.T) {
	t.Parallel()

	want := "The Project Gutenberg eBook of "
	cmd := exec.Command("./cchead", "-c", "31", "test.txt")

	got, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	} else {
		if !cmp.Equal(want, string(got)) {
			t.Errorf("\nwant: \n%s \n\ngot: \n%s", want, string(got))
		}

	}
}

func TestMainStep3_2(t *testing.T) {
	t.Parallel()

	want := "Hello, World\n"
	cmd := exec.Command("./cchead", "-c", "30", "test2.txt")

	got, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	} else {
		if !cmp.Equal(want, string(got)) {
			t.Errorf("\nwant: \n%s \n\ngot: \n%s", want, string(got))
		}

	}
}

func TestMainStep4(t *testing.T) {
	t.Parallel()

	want := "==> test2.txt <==\nHello, World\n==> test.txt <==\nThe Project Gutenberg eBook of The Art of War\r\n    \r\nThis ebook is for the use of anyone anywhere in the United States and\r\nmost other parts of the world at no cost and with almost no restrictions\r\nwhatsoever. You may copy it, give it away or re-use it under the terms\r\nof the Project Gutenberg License included with this ebook or online\r\nat www.gutenberg.org. If you are not located in the United States,\r\nyou will have to check the laws of the country where you are located\r\nbefore using this eBook.\r\n"
	cmd := exec.Command("./cchead", "-n", "9", "test2.txt", "test.txt")

	got, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	} else {
		if !cmp.Equal(want, string(got)) {
			t.Errorf("\nwant: \n%s \n\ngot: \n%s", want, string(got))
		}

	}
}
