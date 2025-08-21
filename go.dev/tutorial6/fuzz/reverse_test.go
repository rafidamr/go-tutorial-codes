package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		inp    string
		expect string
	}{
		{"MyString", "gnirtSyM"},
		{"Hello, World", "dlroW ,olleH"},
		{" ", " "},
	}
	for _, tc := range testcases {
		r, err := Reverse(tc.inp)
		if err != nil {
			t.Skip()
		}
		if r != tc.expect {
			t.Errorf("Get %q, expect %q", r, tc.expect)
		}
	}
}

func FuzzReverse(t *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		t.Add(tc)
	}
	t.Fuzz(func(t *testing.T, ori string) {
		r, err := Reverse(ori)
		if err != nil {
			t.Skip()
		}
		rr, err := Reverse(r)
		if err != nil {
			t.Skip()
		}
		t.Logf("Runes number ori=%d r=%d rr=%d",
			utf8.RuneCountInString(ori),
			utf8.RuneCountInString(r),
			utf8.RuneCountInString(rr))
		if rr != ori {
			t.Errorf("%q and %q are not the same", ori, rr)
		}
		if utf8.ValidString(ori) && !utf8.ValidString(r) {
			t.Errorf("%q is not a valid utf8-encoded string", r)
		}
	})
}
