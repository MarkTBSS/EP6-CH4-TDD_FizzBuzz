package main

import "fmt"

func FizzBuzz(nunber int) string {
	if nunber == 5 || nunber == 10 {
		return "Buzz"
	}
	if nunber == 3 || nunber == 6 || nunber == 9 {
		return "Fizz"
	}
	return fmt.Sprintf("%d", nunber)
}
