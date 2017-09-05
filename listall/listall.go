// -dir only working properly if 2nd arg (or not preceded by non flag args?)
// seems like it would work great if supply necessary ars with flags only.
// can default filetype as `*`

package main

import (
	"fmt"
	"io/ioutil"
	//	"letsTrySomeGo/readline"
	"flag"
	"log"
	"os"
)

func main() {
	dirname := flag.String("dir", "./", "directory not found :(")
	flag.Parse()

	files, err := ioutil.ReadDir(*dirname)
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

	//	if len(args) > 0 {
	//		dir := args[0]
	//		fmt.Println(dir)
	//	}
	if len(args) > 0 {
		fileType := args[0]
		fmt.Println(fileType)
	}
}
