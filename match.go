// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args[1]
	l := strings.Split(input, " ")
	fmt.Println(len(l))
}
