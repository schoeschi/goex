package main

import (
	"fmt"
	"strconv"
)

const (
	Lower = 1
	Upper = 30
)

func main() {
	for i := Lower; i < Upper; i++ {
		word := strconv.Itoa(i)

		if i%3 == 0 && i%5 == 0 {
			word = "FizzBuzz"
		} else if i%5 == 0 {
			word = "Buzz"
		} else if i%3 == 0 {
			word = "Fizz"
		}

		fmt.Printf("%d: %s\n", i, word)
	}
}
