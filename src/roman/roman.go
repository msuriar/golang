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

func init() {
  numerals = append(numerals, Numeral{10, string("X")})
  numerals = append(numerals, Numeral{9, "IX"})
  numerals = append(numerals, Numeral{5, string("V")})
  numerals = append(numerals, Numeral{4, "IV"})
  numerals = append(numerals, Numeral{1, string("I")})
}

func DecimalToRoman(x int) (s string) {
	var outputs []string
	for x > 0 {
    for _,r := range(numerals) {
      for x >= r.dec {
        outputs = append(outputs, r.rom)
        x -= r.dec 
      }
    }
	}
	s = strings.Join(outputs, "")
	return
}
