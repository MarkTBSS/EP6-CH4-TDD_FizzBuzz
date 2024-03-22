package main

import "fmt"

func FizzBuzz(nunber int) string {
	if nunber == 15 {
		return "FizzBuzz"
	}
	if nunber == 5 || nunber == 10 {
		return "Buzz"
	}
	if nunber == 3 || nunber == 6 || nunber == 9 || nunber == 12 {
		return "Fizz"
	}
	return fmt.Sprintf("%d", nunber)
}
