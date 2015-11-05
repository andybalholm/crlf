// The crlf package helps in dealing with files that have DOS-style CR/LF line
// endings.
package crlf

import (
	"io"

	"golang.org/x/text/transform"
)

// Normalize takes CRLF, CR, or LF line endings in src, and converts them
// to LF in dst.
type Normalize struct {
	prev byte
}

func (n *Normalize) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for nDst < len(dst) && nSrc < len(src) {
		c := src[nSrc]
		switch c {
		case '\r':
			dst[nDst] = '\n'
		case '\n':
			if n.prev == '\r' {
				nSrc++
				n.prev = c
				continue
			}
			dst[nDst] = '\n'
		default:
			dst[nDst] = c
		}
		n.prev = c
		nDst++
		nSrc++
	}
	if nSrc < len(src) {
		err = transform.ErrShortDst
	}
	return
}

func (n *Normalize) Reset() {
	n.prev = 0
}

// ToCRLF takes LF line endings in src, and converts them to CRLF in dst.
type ToCRLF struct{}

func (ToCRLF) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for nDst < len(dst) && nSrc < len(src) {
		if c := src[nSrc]; c == '\n' {
			if nDst+1 == len(dst) {
				break
			}
			dst[nDst] = '\r'
			dst[nDst+1] = '\n'
			nSrc++
			nDst += 2
		} else {
			dst[nDst] = c
			nSrc++
			nDst++
		}
	}
	if nSrc < len(src) {
		err = transform.ErrShortDst
	}
	return
}

func (ToCRLF) Reset() {}

// NewReader returns an io.Reader that converts CR or CRLF line endings to LF.
func NewReader(r io.Reader) io.Reader {
	return transform.NewReader(r, new(Normalize))
}

// NewWriter returns an io.Writer that converts LF line endings to CRLF.
func NewWriter(w io.Writer) io.Writer {
	return transform.NewWriter(w, ToCRLF{})
}
