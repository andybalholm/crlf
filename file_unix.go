// +build !windows

package crlf

import (
	"io"
	"os"
)

// Open opens a text file for reading, with platform-appropriate line ending
// conversion.
func Open(name string) (io.ReadCloser, error) {
	return os.Open(name)
}

// Create opens a text file for writing, with platform-appropriate line ending
// conversion.
func Create(name string) (io.WriteCloser, error) {
	return os.Create(name)
}
