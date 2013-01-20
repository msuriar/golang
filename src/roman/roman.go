package roman

import (
	"strings"
)

// import "fmt"

type Numeral struct {
	dec int
	rom string
}

var numerals []Numeral

var characters map[rune]int

func init() {
	numerals = append(numerals, Numeral{1000, string("M")})
	numerals = append(numerals, Numeral{900, string("CM")})
	numerals = append(numerals, Numeral{500, string("D")})
	numerals = append(numerals, Numeral{400, string("CD")})
	numerals = append(numerals, Numeral{100, string("C")})
	numerals = append(numerals, Numeral{90, string("XC")})
	numerals = append(numerals, Numeral{50, string("L")})
	numerals = append(numerals, Numeral{40, string("XL")})
	numerals = append(numerals, Numeral{10, string("X")})
	numerals = append(numerals, Numeral{9, "IX"})
	numerals = append(numerals, Numeral{5, string("V")})
	numerals = append(numerals, Numeral{4, "IV"})
	numerals = append(numerals, Numeral{1, string("I")})

  characters = make(map[rune]int)
  characters['I'] = 1
  characters['V'] = 5
  characters['X'] = 10
  characters['L'] = 50
  characters['C'] = 100
  characters['D'] = 500
  characters['M'] = 1000
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
  for i := 0 ; i < len(s) ; i++ {
    current,next := characters[rune(s[i])],0

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
