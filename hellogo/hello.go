// NOTE: I'm definitely retarded.  Exiting the program from greet() makes it work as expected
// UNLESS: you hit enter without typing anything.  a single character routes through the first part of the if{}
// > 1 char gives you the else, but no chars just returns without running greet.

// NOTE: just moved this project from arch linux to windows10... this bug
// no longer takes effect, and the program seems to work as expected. Interesting.

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
			greet(string(len(r))) // this is only reached if len(r) == 1, not if len(r) < 1.
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
