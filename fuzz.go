// +build gofuzz

package crlf

import "golang.org/x/text/transform"

func Fuzz(data []byte) int {
	_, _, err := transform.Bytes(new(Normalize), data)
	if err != nil {
		panic(err)
	}

	_, _, err = transform.Bytes(ToCRLF{}, data)
	if err != nil {
		panic(err)
	}

	return 0
}
