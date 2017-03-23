package golangttdtextcalctravis

import (
  "testing" // import the testing capability of go
)

func invoke_text_to_number(input string, expected_output int, t *testing.T) {
  output := Text_to_number(input)
  if  output != expected_output {
    t.Errorf("FAILED: text_to_number(%s) returned %d, expected %d", input, output, expected_output)
  } else {
    t.Logf("PASSED: The output was correctly %d", output)
  }
}

func TestZero(t *testing.T) {
  input := "zero"
  expected_output:= 0;
  invoke_text_to_number(input, expected_output, t)
}

func TestOne(t *testing.T) {
  input := "one"
  expected_output:= 1;
  invoke_text_to_number(input, expected_output, t)
}

func TestFive(t *testing.T) {
  input := "five"
  expected_output:= 5;
  invoke_text_to_number(input, expected_output, t)
}

/* // example of failing test
func TestSix(t *testing.T) {
  input := "six"
  expected_output:= 6;
  invoke_text_to_number(input, expected_output, t)
}
*/
