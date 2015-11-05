package crlf

import (
	"io"
	"os"
)

type readCloser struct {
	io.Reader
	io.Closer
}

type writeCloser struct {
	io.Writer
	io.Closer
}

// Open opens a text file for reading, with platform-appropriate line ending
// conversion.
func Open(name string) (io.ReadCloser, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return readCloser{
		Reader: NewReader(f),
		Closer: f,
	}, nil
}

// Create opens a text file for writing, with platform-appropriate line ending
// conversion.
func Create(name string) (io.WriteCloser, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	return writeCloser{
		Writer: NewWriter(f),
		Closer: f,
	}, nil
}
