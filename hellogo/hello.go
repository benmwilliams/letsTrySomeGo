// What I learned: this silly little project frustrated me for a long time.
// For really no reason other than not understanding why, I wanted to be able
// to have an enter keypress just log and empty response and close out the
// program. It never worked on linux, but then when I tried on windows, it
// worked perfectly. Now the mystery is why it works on windows, and the truth
// is that is was working on linux all the time. The readline package trims the
// newline character from the input, so in the console you get a new line, but
// the program doesn't terminate. Once I removed the line that trimmed `\n`,
// pressing enter would terminate the program with the greeting adressed to
// "newline binary character". I now realize this is the only reasonable
// behavior. My frustration was not from this code not doing what I wanted,
// but from my wanting something illogical. I don't think I even need the
// readline package to do this, but I'm ready to close this and move on to
// something else. This should really just be like 15 lines of code.
// buuuut... why *does* it work the way I wanted on windows?

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
		if len(r) < 1 || line == "\n" {
			greet(string(len(r))) // this would never exec bc readline trims \n
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
