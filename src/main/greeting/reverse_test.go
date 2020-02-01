package greeting_test

import (
	"github.com/barbarum/go-test/src/main/greeting"
	"log"
	"os"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello World", "dlroW olleH"},
		{"", ""},
		{"i", "i"},
	}

	for _, c := range cases {
		got := greeting.Reverse(c.in)

		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFileOpen(t *testing.T) {
	_, err := os.Open("./testdata/dummy.txt")

	if err != nil {
		log.Fatal(err)
	}

}
