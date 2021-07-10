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
	if strings.EqualFold(input, "") {
		fmt.Println("0")
	} else {
		l := strings.Split(input, " ")
		fmt.Println(len(l))
	}
}
