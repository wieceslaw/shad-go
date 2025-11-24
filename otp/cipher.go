//go:build !solution

package otp

import (
	"io"
)

type xorReader struct {
	r    io.Reader
	prng io.Reader
}

func (x *xorReader) Read(p []byte) (n int, err error) {
	n, err = x.r.Read(p)
	if n == 0 && err != nil {
		return 0, err
	}

	arr := make([]byte, n)

	_, _ = x.prng.Read(arr)

	for i := range arr {
		p[i] ^= arr[i]
	}

	return n, err
}

func NewReader(r io.Reader, prng io.Reader) io.Reader {
	return &xorReader{r, prng}
}

type xorWriter struct {
	w    io.Writer
	prng io.Reader
}

func (w *xorWriter) Write(p []byte) (int, error) {
	arr := make([]byte, len(p))
	_, _ = w.prng.Read(arr)

	for i := range p {
		arr[i] ^= p[i]
	}

	return w.w.Write(arr)
}

func NewWriter(w io.Writer, prng io.Reader) io.Writer {
	return &xorWriter{w, prng}
}
