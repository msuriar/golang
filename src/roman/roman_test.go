package roman

import "testing"

func TestDecimalToRoman(t *testing.T) {
  const in,out = 1,string('I')
  if x:= DecimalToRoman(in); x != out {
    t.Errorf("DecimalToRoman(%v) = %v, want %v", in, x, out)
  }
}
