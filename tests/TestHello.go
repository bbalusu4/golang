package main

import "testing"
//testing Hello World program

func TestHello(t *testing.T){
  testFunc := main() //main function does not return any value
  testString := "Hello World"

  if testFunc != testString {
	  t.errorf("testFunc %q testString %q", testFunc, testString)
  }
}
