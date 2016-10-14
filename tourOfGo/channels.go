package main

import "fmt"

func sum(s []int, c chan int) { // sum() takes an array of ints and an int channel
	sum := 0              // var sum int = 0
	for _, v := range s { // ?????
		sum += v // sum = sum + v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0} // var s int = [7, 2, 8, -9, 4, 0]

	c := make(chan int)     // make???
	go sum(s[:len(s)/2], c) // run sum(), passing
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
