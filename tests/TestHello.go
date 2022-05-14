package main

import "testing"
//testing Hello World program

func TestHello(t *testing.T){
 // testFunc := main() //main function does not return any value
  testString := "Hello World"
  testFunc := Hello() //new function created to have retrun value and called to print from main func

  if testFunc != testString {
	  t.errorf("testFunc %q testString %q", testFunc, testString)
  }
}
