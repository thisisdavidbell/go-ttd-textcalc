package gottdtextcalc

import (
  "testing" // import the testing capability of go
)

func TestZero(t *testing.T) {
//same as following  var one string = "one"
//  one := "one";
  input := "zero"
  expected_output:= 0;

  output := text_to_number(input)
  t.Logf("Logf: the output is %d", output)
  if  output != expected_output {
    t.Error("text_to_number(%v) returned %v, expected %v", input, output, expected_output)
  } else {
    t.Log("TEST PASSED")
  }
}

func TestOne(t *testing.T) {
//same as following  var one string = "one"
//  one := "one";
  input := "one"
  expected_output:= 1;

  output := text_to_number(input)
  t.Logf("Logf: the output is %d", output)
  if  output != expected_output {
    t.Error("text_to_number(%v) returned %v, expected %v", input, output, expected_output)
  } else {
    t.Log("TEST PASSED")
  }
}

func TestSix(t *testing.T) {
//same as following  var one string = "one"
//  one := "one";
  input := "six"
  expected_output:= 6;

  output := text_to_number(input)
  t.Logf("Logf: the output is %d", output)
  if  output != expected_output {
    t.Errorf("text_to_number(%s) returned %d, expected %d", input, output, expected_output)
  } else {
    t.Log("TEST PASSED")
  }
}
