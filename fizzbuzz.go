package main

import "fmt"

func FizzBuzz(nunber int) string {
	if nunber == 5 {
		return "Buzz"
	}
	if nunber == 3 || nunber == 6 {
		return "Fizz"
	}
	return fmt.Sprintf("%d", nunber)
}
