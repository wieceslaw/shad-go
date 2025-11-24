//go:build !solution

package varfmt

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func readNumber(str []rune) (n int, idx int, err error) {
	if len(str) == 0 || !unicode.IsDigit(str[0]) {
		err = errors.New("not a number")
		return
	}

	idx = 0
	n = 0
	err = nil
	for idx < len(str) && unicode.IsDigit(str[idx]) {
		d := int(str[idx]) - '0'
		n = n*10 + d
		idx++
	}
	return
}

func readFormatNumber(str []rune) (n int, idx int, err error) {
	if len(str) == 0 || str[0] != '{' {
		err = errors.New("is not a format string")
		return
	}

	n, idx, err = readNumber(str[1:])
	if err != nil {
		return
	}

	idx++
	if str[idx] != '}' {
		err = errors.New("is not a format string")
		return
	}

	return
}

func readFormat(str []rune) (idx int, err error) {
	if len(str) < 2 || str[0] != '{' || str[1] != '}' {
		err = errors.New("not a format")
		return
	}

	idx += 1
	return
}

func Sprintf(format string, args ...interface{}) string {
	var b strings.Builder
	runes := []rune(format)
	formats := 0
	for i := 0; i < len(runes); i++ {
		n, idx, err := readFormatNumber(runes[i:])
		if err == nil {
			if n >= len(args) {
				panic("argument too big")
			}
			arg := fmt.Sprintf("%v", args[n])
			b.WriteString(arg)
			i += idx
			formats++
			continue
		}

		idx, err = readFormat(runes[i:])
		if err == nil {
			if formats >= len(args) {
				panic("argument too big")
			}
			arg := fmt.Sprintf("%v", args[formats])
			b.WriteString(arg)
			i += idx
			formats++
			continue
		}

		b.WriteRune(runes[i])
	}
	return b.String()
}
