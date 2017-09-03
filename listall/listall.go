package main

import (
	"fmt"
	"io/ioutil"
	//	"letsTrySomeGo/readline"
	"log"
	"os"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
	printArgs()
}

func printArgs() {
	args := os.Args[1:]
	fmt.Println(args)
}
