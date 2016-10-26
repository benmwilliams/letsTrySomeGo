// NOTE: I'm definitely retarded.  Exiting the program from greet() makes it work as expected
// UNLESS: you hit enter without typing anything.  a single character routes through the first part of the if{}
// > 1 char gives you the else, but no chars just returns without running greet.

package main

import (
	"fmt"
	"os"

	"github.com/benmwilliams/letsTrySomeGo/readline"
)

func main() {
	name := ""
	fmt.Printf("What is your name? \n")

	readline.ReadLine(os.Stdin, func(line string) {
		r := []rune(line)
		if len(r) <= 1 || line == "" {
			greet("found me")
		} else {
			name = line
			greet(name)
		}
	})
}

func greet(n string) {
	fmt.Printf("Hello, %v \nHave a nice day! \n", n)
	os.Exit(0)
}
