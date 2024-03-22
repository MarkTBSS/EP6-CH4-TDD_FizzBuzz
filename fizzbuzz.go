package main

import "fmt"

func FizzBuzz(nunber int) string {
	if isFizzBuzz(nunber) {
		return "FizzBuzz"
	}
	if isBuzz(nunber) {
		return "Buzz"
	}
	if isFizz(nunber) {
		return "Fizz"
	}
	return fmt.Sprintf("%d", nunber)
}

func isFizz(nunber int) bool {
	return nunber%3 == 0
}

func isBuzz(nunber int) bool {
	return nunber%5 == 0
}

func isFizzBuzz(nunber int) bool {
	return nunber%15 == 0
}
