package operator_test

import (
	"fmt"
	operator "operator/operator"
)

func Example() {
	operator.Prints("hello")

	// Output:
	// hello
}

func Example_greet() {
	n := operator.Normalizer{}
	fmt.Println(n.Greet("Franklin"))

	n.Formal = true
	fmt.Println(n.Greet("Wilson"))

	// Output:
	// Hello Franklin
	// Hello Mr Wilson
}
