// NOTE: readline now works. Program will not exit, I don't know why. It just keeps reading from stdin, and nothing else happens,
// even if I try to return from the function, or os.Exit(0) from within the function. Pretty sure I'm just retarded.

package main

import (
	"fmt"
	"os"

	"github.com/benmwilliams/letsTrySomeGo/readline"
)

func main() {
	fmt.Printf("What is your name? \n")
	readline.ReadLine(os.Stdin, func(name string) {
		if name == "\r\n" {
			fmt.Printf("found me")
		} else {
			fmt.Printf("Hello, %v \n", name)
			fmt.Printf("Have a nice day! \n")
			fmt.Printf("name is %v . \n", name)
		}
	})
	fmt.Printf("not in readline")
}
