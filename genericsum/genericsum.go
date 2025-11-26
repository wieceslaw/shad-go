//go:build !solution

package genericsum

import (
	"cmp"
	"sync"
)

func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func SortSlice[T cmp.Ordered](a []T) {
	for i := range a {
		for j := i; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func MapsEqual[K comparable, V comparable](a, b map[K]V) bool {
	for k, v1 := range a {
		if v2, ok := b[k]; !ok || v2 != v1 {
			return false
		}
	}
	for k, v1 := range b {
		if v2, ok := a[k]; !ok || v2 != v1 {
			return false
		}
	}
	return true
}

func SliceContains[T comparable](s []T, v T) bool {
	for _, el := range s {
		if el == v {
			return true
		}
	}
	return false
}

func MergeChans[T any](chs ...<-chan T) <-chan T {
	resCh := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(chs))
	for _, ch := range chs {
		go func() {
			for el := range ch {
				resCh <- el
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(resCh)
	}()
	return resCh
}

type Numeric interface {
	~int | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int8 | ~int16 | ~int32 | ~int64 | ~complex128 | ~complex64 | ~float32 | ~float64
}

func IsHermitianMatrix[T Numeric](m [][]T) bool {
	if len(m) == 0 {
		return true
	}
	if len(m) != len(m[0]) {
		return false
	}

	for i := range m {
		for j := range m[i] {
		}
	}
	// equals to transposed matrix with all elements conjugated (real part is intact, while imaginary part is inverted)
	panic("implement me")
}
