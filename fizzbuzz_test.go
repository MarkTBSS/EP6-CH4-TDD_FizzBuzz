package main

import (
	"testing"
)

type Case struct {
	name   string
	number int
	want   string
}

func TestFizzBuzz(t *testing.T) {
	testcase_1 := Case{name: "FizzBuzz(1)", number: 1, want: "1"}
	testcase_3 := Case{name: "FizzBuzz(3)", number: 3, want: "Fizz"}
	testcase_5 := Case{name: "FizzBuzz(5)", number: 5, want: "Buzz"}
	testcase_15 := Case{name: "FizzBuzz(15)", number: 15, want: "FizzBuzz"}
	tests := []Case{
		testcase_1,
		testcase_3,
		testcase_5,
		testcase_15,
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			got := FizzBuzz(testcase.number)
			if got != testcase.want {
				t.Errorf("FizzBuzz(%d) = %s; want %s", testcase.number, got, testcase.want)
			}
		})
	}
}
