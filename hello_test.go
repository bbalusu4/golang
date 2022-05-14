package HelloWorld

import "testing"

//testing Hello World program

func TestHello(t *testing.T) {
	// testFunc := main() //main function does not return any value
	testM := Hello() //new function created to have retrun value and called to print from main func
	testS := "Hello World!"

	if testM != testS {
		t.Errorf("testM %q testS %q", testM, testS)
	}
}
