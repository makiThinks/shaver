package test

import (
	"testing"

	"maki.io/demo/shaver/util"
)

type TestData struct {
	in, want string
}

func TestReverseString(t *testing.T) {

	cases := []TestData{
		{"hello", "olleh"},
		{"1", "1"},
		{"", ""},
		{"世界你好", "好你界世"},
	}

	for _, c := range cases {
		got := util.ReverseString(c.in)
		if got == c.want {
			t.Logf("correct: %q=%q", c.want, got)
			continue
		} else {
			t.Errorf("ReverseString(%q)==%q, want %q", c.in, got, c.want)
		}
	}
}
