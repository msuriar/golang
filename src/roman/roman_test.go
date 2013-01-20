package roman

import "testing"

var tests map[int]string

func init() {
	tests = map[int]string{
		1:    "I",
		3:    "III",
		4:    "IV",
		5:    "V",
		8:    "VIII",
		9:    "IX",
		10:   "X",
		13:   "XIII",
		14:   "XIV",
		15:   "XV",
		16:   "XVI",
		19:   "XIX",
		20:   "XX",
		90:   "XC",
		100:  "C",
		1900: "MCM",
		1990: "MCMXC",
	}
}

func TestRomanToDecimal(t *testing.T) {
	for out, in := range tests {
		if x := RomanToDecimal(in); x != out {
			t.Errorf("RomanToDecimal(%v) = %v, want %v", in, x, out)
		}
	}
}

func TestDecimalToRoman(t *testing.T) {
	for in, out := range tests {
		if x := DecimalToRoman(in); x != out {
			t.Errorf("DecimalToRoman(%v) = %v, want %v", in, x, out)
		}
	}
}
