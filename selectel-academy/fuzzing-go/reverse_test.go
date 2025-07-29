package fuzzinggo

import (
	"testing"
	"unicode/utf8"
)

func TestReverseSimple(t *testing.T) {
	expected := "!dlroW ,olleH"
	got, error := Reverse("Hello, World!")
	if error != nil {
		t.Error(error)
	}

	if got != expected {
		t.Errorf("actual: %q, expected: %q", got, expected)
	}
}

func FuzzReverse(f *testing.F) {
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		doubleRev, _ := Reverse(rev)

		if err1 != nil {
			t.Error(err1)
		}

		if orig != doubleRev {
			t.Errorf("Before: %q, after %q", orig, doubleRev)
		}

		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produces invalid UTF-8 string %q", rev)
		}
	})
}
