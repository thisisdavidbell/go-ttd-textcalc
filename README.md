# go-ttd-textcalc
A tutorial which uses Test Driven Development with Go to build a text based calculator.

This project is the root repo, and main README, of a tutorial to try out the built in testing capability of Go, using TDD to build a text based calculator. i.e. you can pass in three plus five and you will getback the answer eight. However, for now, lets just write a tool which allows you to enter a string from 0 to 20, and returns it as a number, e.g. text_to_number("three") returns 3.


# Creating the workspace
* Create a directory structure with the main top level dir for a workspace, and the path to reflect the repo(s), eg. `go-ttd-text-calc/github.com/thisisdavidbell`
* clone this repo into that dir to produce dir structure `go-ttd-text-calc/github.com/thisisdavidbell/go-ttd-textcalc`.
* Now we need to create the tests to check the text_to_number() method works correctly. We of course alreay have the finished file in test_textcalc.go


# Useful links
* Intro to Go: https://tour.golang.org
