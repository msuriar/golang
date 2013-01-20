package roman

import "testing"

func TestDecimalToRoman(t *testing.T) {
  tests := map[int]string{
    1: string('I'),
    3: "III",
    4: "IV",
    5: string('V'),
    8: "VIII",
    9: "IX",
    10: "X",
  }
  for in,out := range tests {
    if x:= DecimalToRoman(in); x != out {
      t.Errorf("DecimalToRoman(%v) = %v, want %v", in, x, out)
    }
}
}
