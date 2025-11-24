//go:build !solution

package speller

import (
	"strings"
)

var numbers = [...]string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}
var decimals = [...]string{
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}
var orders = [...]string{"", "thousand", "million", "billion", "trillion"}

func spellHundreds(n int64) string {
	var b strings.Builder

	if n >= 100 {
		b.WriteRune(' ')
		b.WriteString(numbers[n/100])
		b.WriteRune(' ')
		b.WriteString("hundred")
		n = n % 100
	}
	if n > 0 {
		if n < 20 {
			b.WriteRune(' ')
			b.WriteString(numbers[n])
		} else {
			b.WriteRune(' ')
			b.WriteString(decimals[(n/10)-2])
			if n%10 != 0 {
				b.WriteRune('-')
				b.WriteString(numbers[n%10])
			}
		}
	}

	return b.String()
}

func Spell(n int64) string {
	if n == 0 {
		return "zero"
	} else if n == 1000000 {
		return "one million"
	} else if n == 1000000000 {
		return "one billion"
	}

	var result string
	var negative bool
	if n < 0 {
		negative = true
		n = -n
	}
	order := 0
	for n > 0 {
		s := spellHundreds(n % 1000)
		result = s + " " + orders[order] + result
		n = n / 1000
		order++
	}

	if negative {
		return "minus " + strings.Trim(result, " ")
	}
	return strings.Trim(result, " ")
}
