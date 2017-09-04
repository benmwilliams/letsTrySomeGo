package main

import (
	"fmt"
	"io/ioutil"
	//	"letsTrySomeGo/readline"
	//	"flag"
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
	getArgs()
}

func getArgs() {
	args := os.Args[1:]

	if len(args) > 0 {
		dir := args[0]
		fmt.Println(dir)
	}
	if len(args) > 1 {
		fileType := args[1]
		fmt.Println(fileType)
	}
}
