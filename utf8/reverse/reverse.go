//go:build !solution

package reverse

import (
	"strings"
	"unicode/utf8"
)

func Reverse(input string) string {
	runes := []rune(input)
	var b strings.Builder
	b.Grow(len(input))
	for i := len(runes) - 1; i >= 0; i-- {
		rn := runes[i]
		if utf8.ValidRune(rn) {
			b.WriteRune(rn)
		} else {
			b.WriteRune(utf8.RuneError)
		}
	}

	return b.String()
}
