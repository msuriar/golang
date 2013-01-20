package roman

import (
	"strings"
)

type Numeral struct {
	dec int
	rom string
}

var characters = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var numerals = []Numeral{
	0:  Numeral{1000, "M"},
	1:  Numeral{900, "CM"},
	2:  Numeral{500, "D"},
	3:  Numeral{400, "CD"},
	4:  Numeral{100, "C"},
	5:  Numeral{90, "XC"},
	6:  Numeral{50, "L"},
	7:  Numeral{40, "XL"},
	8:  Numeral{10, "X"},
	9:  Numeral{9, "IX"},
	10: Numeral{5, "V"},
	11: Numeral{4, "IV"},
	12: Numeral{1, "I"},
}

func DecimalToRoman(x int) (s string) {
	var outputs []string
	for x > 0 {
		for _, r := range numerals {
			for x >= r.dec {
				outputs = append(outputs, r.rom)
				x -= r.dec
			}
		}
	}
	s = strings.Join(outputs, "")
	return
}

func RomanToDecimal(s string) (x int) {
	x = 0
	for i := 0; i < len(s); i++ {
		current, next := characters[rune(s[i])], 0

		if i+1 < len(s) {
			next = characters[rune(s[i+1])]
		}

		if current < next {
			x += next - current
			i++
		} else {
			x += current
		}
	}
	return x
}
