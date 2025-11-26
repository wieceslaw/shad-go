//go:build !solution

package externalsort

import (
	"bufio"
	"container/heap"
	"errors"
	"io"
	"os"
	"sort"
	"strings"
)

type LineReaderImpl struct {
	r io.Reader
}

func (lr *LineReaderImpl) ReadLine() (string, error) {
	r := lr.r
	var sb strings.Builder
	buf := make([]byte, 1)

	for {
		n, err := r.Read(buf)

		if n > 0 {
			b := buf[0]
			if b == '\n' {
				break
			}
			sb.WriteByte(b)
		}

		if err != nil {
			if err == io.EOF {
				if sb.Len() == 0 {
					return "", io.EOF
				}
				break
			}
			return "", err
		}
	}
	return sb.String(), nil
}

type LineWriterImpl struct {
	w io.Writer
}

func (lw *LineWriterImpl) Write(l string) error {
	_, err := lw.w.Write([]byte(l))
	if err != nil {
		return err
	}
	_, err = lw.w.Write([]byte{'\n'})
	return err
}

func NewReader(r io.Reader) LineReader {
	return &LineReaderImpl{r}
}

func NewWriter(w io.Writer) LineWriter {
	return &LineWriterImpl{w}
}

type pair struct {
	s string
	i int
}

type stringHeap []pair

func (sh stringHeap) Len() int {
	return len(sh)
}

func (sh stringHeap) Less(i, j int) bool {
	return sh[i].s < sh[j].s
}

func (sh stringHeap) Swap(i, j int) {
	sh[i], sh[j] = sh[j], sh[i]
}

func (sh *stringHeap) Push(x any) {
	*sh = append(*sh, x.(pair))
}

func (sh *stringHeap) Pop() any {
	n := len(*sh)
	x := (*sh)[n-1]
	*sh = (*sh)[0 : n-1]
	return x
}

func Merge(w LineWriter, readers ...LineReader) error {
	h := &stringHeap{}
	for i, r := range readers {
		line, err := r.ReadLine()
		switch {
		case errors.Is(err, io.EOF):
			continue
		case err != nil:
			return err
		}
		heap.Push(h, pair{line, i})
	}
	for h.Len() != 0 {
		top := heap.Pop(h).(pair)
		w.Write(top.s)
		r := readers[top.i]
		line, err := r.ReadLine()
		switch {
		case err == nil:
			heap.Push(h, pair{line, top.i})
		case errors.Is(err, io.EOF):
			continue
		default:
			return err
		}
	}
	return nil
}

func Sort(w io.Writer, in ...string) error {
	var readers []LineReader
	for _, file := range in {
		if err := sortFile(file); err != nil {
			return err
		}

		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()
		// sort each file
		r := NewReader(bufio.NewReader((f)))
		readers = append(readers, r)
	}

	return Merge(NewWriter(w), readers...)
}

func sortFile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := NewReader(bufio.NewReader(f))
	var lines []string
	for {
		line, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		lines = append(lines, line)
	}

	sort.Strings(lines)

	f, err = os.Create(file)
	if err != nil {
		return err
	}
	bw := bufio.NewWriter(f)
	w := NewWriter(bw)
	for _, line := range lines {
		if err := w.Write(line); err != nil {
			return err
		}
	}
	if err := bw.Flush(); err != nil {
		return err
	}

	return nil
}
