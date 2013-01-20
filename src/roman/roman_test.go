package roman

import "testing"

func TestDecimalToRoman(t *testing.T) {
	tests := map[int]string{
		1:   string('I'),
		3:   "III",
		4:   "IV",
		5:   string('V'),
		8:   "VIII",
		9:   "IX",
		10:  "X",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		19:  "XIX",
		20:  "XX",
		90:  "XC",
		100: "C",
	}
	for in, out := range tests {
		if x := DecimalToRoman(in); x != out {
			t.Errorf("DecimalToRoman(%v) = %v, want %v", in, x, out)
		}
	}
}
