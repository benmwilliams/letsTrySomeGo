// TODO: moved this file, changed the function of the program to ask for a name and then greet.
// the original binary is still in go/bin, as well as the newly created hellogo.
// this shows that the binary is created based on the dir name, and not the name of the source file
// also there is a binary in this directory, and Im not sure if it's a result of a bad config of an errant build step
// running both `hello` and `hellogo` work from any dir, so GOPATH and GOBIN must be correctly set.

package main

import "fmt"

func main() {
	fmt.Printf("What is your name? \n")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hello, %v \n", name)
}
