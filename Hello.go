package HelloWorld

//simple Hello World program
import "fmt"

func main() {
	fmt.Println(Hello())
}

func Hello() string {
	return "Hello World!"
}

func Hello_value(name string) string{
	return "Hello, i am " + name
}
