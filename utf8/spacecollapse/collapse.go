//go:build !solution

package spacecollapse

import (
	"strings"
	"unicode"
)

func CollapseSpaces(input string) string {
	var b strings.Builder
	b.Grow(len(input))
	flag := false
	for _, rn := range input {
		if unicode.IsSpace(rn) {
			flag = true
		} else {
			if flag {
				b.WriteString(" ")
				flag = false
			}
			b.WriteRune(rn)
		}
	}
	if flag {
		b.WriteString(" ")
	}
	return b.String()
}
