package main

import "fmt"

func FizzBuzz(nunber int) string {
	if nunber%15 == 0 {
		return "FizzBuzz"
	}
	if nunber%5 == 0 {
		return "Buzz"
	}
	if nunber%3 == 0 {
		return "Fizz"
	}
	return fmt.Sprintf("%d", nunber)
}
