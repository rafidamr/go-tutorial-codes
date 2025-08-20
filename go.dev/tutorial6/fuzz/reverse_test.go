package main

import "testing"

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
		r := Reverse(tc.inp)
		if r != tc.expect {
			t.Errorf("Get %q, expect %q", r, tc.expect)
		}
	}
}
