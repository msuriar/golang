package roman

import (
  "strings"
)

// import "fmt"

func DecimalToRoman(x int) (s string) {
  var outputs []string
  for x > 0 {
    switch {
      case x >= 5:
        outputs = append(outputs, string("V"))
        x -= 5
      case x >= 4:
        outputs = append(outputs, string("IV"))
        x -= 4
      default:
        outputs = append(outputs, string("I"))
        x--
    }
  }
  s = strings.Join(outputs, "")
  return
}
