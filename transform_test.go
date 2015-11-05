package crlf

import (
	"testing"

	"golang.org/x/text/transform"
)

func TestNormalize(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{"hello, world\r\n", "hello, world\n"},
		{"hello, world\r", "hello, world\n"},
		{"hello, world\n", "hello, world\n"},
		{"", ""},
		{"\r\n", "\n"},
		{"hello,\r\nworld", "hello,\nworld"},
		{"hello,\rworld", "hello,\nworld"},
		{"hello,\nworld", "hello,\nworld"},
		{"hello,\n\rworld", "hello,\n\nworld"},
		{"hello,\r\n\r\nworld", "hello,\n\nworld"},
	}

	n := new(Normalize)

	for _, c := range testCases {
		got, _, err := transform.String(n, c.in)
		if err != nil {
			t.Errorf("error transforming %q: %v", c.in, err)
			continue
		}
		if got != c.want {
			t.Errorf("transforming %q: got %q, want %q", c.in, got, c.want)
		}
	}
}

func TestToCRLF(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{"hello, world\n", "hello, world\r\n"},
		{"", ""},
		{"\n", "\r\n"},
		{"hello,\nworld", "hello,\r\nworld"},
		{"hello,\n\nworld", "hello,\r\n\r\nworld"},
	}

	for _, c := range testCases {
		got, _, err := transform.String(ToCRLF{}, c.in)
		if err != nil {
			t.Errorf("error transforming %q: %v", c.in, err)
			continue
		}
		if got != c.want {
			t.Errorf("transforming %q: got %q, want %q", c.in, got, c.want)
		}
	}
}
