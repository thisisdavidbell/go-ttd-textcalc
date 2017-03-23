# go-ttd-textcalc
A tutorial which uses Test Driven Development with Go to build a text based calculator.

This project is the root repo, and main README, of a tutorial to try out the built in testing capability of Go, using TDD to build a text based calculator. i.e. you can pass in three plus five and you will getback the answer eight. However, for now, lets just write a tool which allows you to enter a string from 0 to 20, and returns it as a number, e.g. text_to_number("three") returns 3.

Later we add Travis integration to auto build and test it on code changes.

# Creating the workspace
* Create a directory structure with the main top level dir for a workspace, and the path to reflect the repo(s), eg. `go-ttd-text-calc/github.com/thisisdavidbell`
* clone this repo into that dir to produce dir structure `go-ttd-text-calc/github.com/thisisdavidbell/golangttdtextcalc`.
* Now we need to create the tests to check the text_to_number() method works correctly. We of course already have the finished file in [test_textcalc.go](test_textcalc.go)
* Now we run the test: `go test github.com/thisisdavidbell/golangttdtextcalc`. This produces:
```
./textcalc_test.go:8: undefined: text_to_number
FAIL	github.com/thisisdavidbell/golangttdtextcalc [build failed]
```
  * This fails as the implementation of the function text_to_number doesnt exist
* Implement the method. Create a new file containing the method text_to_number in a file with matching name minus the 'test_' and wit the same package name, which by convension is the same as the same as the last part of the import path, i.e. the dir name the live in, so here 'textcalc.go'. You already have an example in this repo: [textcal.go](textcalc.go)
* Now lets run the test again: `go test github.com/thisisdavidbell/golangttdtextcalc`
  * No output other than 'ok' means test passed.
  * You can run with -v to see the t.Logf output.
  * Note: the package is not installed into pkg. Its presumably built every time and thrown away.

# Useful links
* Intro to Go: https://tour.golang.org
