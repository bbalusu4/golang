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

//test with value assignment
func TestAssignValue(t *testing.T) {
	fmt.Println("running in TestAssignValue")
	assertMessage := func(t testing.TB, aValue, aResult string) {
		t.Helper()
		if aValue != aResult {
			t.Errorf("aValue %q aResult %q", aValue, aResult)
		}

	}

	t.Run("Saying Hello to assigned name", func(t *testing.T) {
		aValue := Hello_value("Bala")
		aResult := "Hello, i am Bala"
		assertMessage(t, aValue, aResult)
	})
	t.Run("saying Hello World", func(t *testing.T) {
		aValue := Hello_value("")
		aResult := "Hello, i am world"
		assertMessage(t, aValue, aResult)
	})

}
