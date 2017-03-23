package main

import (
  "github.com/thisisdavidbell/golangttdtextcalctravis"
  "fmt"
)

func main() {
  input := "four"
  output := golangttdtextcalctravis.Text_to_number(input)
  fmt.Printf("text_to_number(%s) from imported package returned %d", input, output)
}
