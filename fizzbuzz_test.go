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
	testcase_2 := Case{name: "FizzBuzz(2)", number: 2, want: "2"}
	testcase_3 := Case{name: "FizzBuzz(3)", number: 3, want: "Fizz"}
	tests := []Case{
		testcase_1,
		testcase_2,
		testcase_3,
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
