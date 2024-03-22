package main

import "fmt"

func FizzBuzz(nunber int) string {
	/* if nunber == 4 {
		return "4"
	} */
	if nunber == 3 {
		return "Fizz"
	}
	/* if nunber == 2 {
		return "2"
	} */
	s := fmt.Sprintf("%d", nunber)
	return s
}
