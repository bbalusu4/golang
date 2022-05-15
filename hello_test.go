package HelloWorld

import (
	"fmt"
	"testing"
)

//testing Hello World program

func TestHello(t *testing.T) {
	// testFunc := main() //main function does not return any value
	testM := Hello() //new function created to have retrun value and called to print from main func
	testS := "Hello World!"
	fmt.Println("running in TestHello")

	if testM != testS {
		t.Errorf("testM %q testS %q", testM, testS)
	}
}

//test with dynamic value
func TestValue(t *testing.T) {
	fmt.Println("running in TestValue")
	valuePassed := Hello_value("testing")
	result := "Hello, i am testing"

	if valuePassed != result {
		t.Errorf("valuePassed %q result %q", valuePassed, result)
	}
}
