package operator

import (
	"fmt"
	"strings"
)

func Prints(s string) {
	fmt.Println(s)
}

func Format(s string) string {
	return strings.ToUpper(s)
}

type Normalizer struct {
	Formal bool
}

func (n *Normalizer) Greet(name string) string {
	if n.Formal {
		return fmt.Sprintf("Hello Mr %s", name)
	}
	return fmt.Sprintf("Hello %s", name)
}
