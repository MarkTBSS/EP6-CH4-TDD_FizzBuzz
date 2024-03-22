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
	testcase_4 := Case{name: "FizzBuzz(4)", number: 4, want: "4"}
	testcase_5 := Case{name: "FizzBuzz(5)", number: 5, want: "Buzz"}
	testcase_6 := Case{name: "FizzBuzz(6)", number: 6, want: "Fizz"}
	testcase_7 := Case{name: "FizzBuzz(7)", number: 7, want: "7"}
	testcase_8 := Case{name: "FizzBuzz(8)", number: 8, want: "8"}
	testcase_9 := Case{name: "FizzBuzz(9)", number: 9, want: "Fizz"}
	testcase_10 := Case{name: "FizzBuzz(10)", number: 10, want: "Buzz"}
	tests := []Case{
		testcase_1,
		testcase_2,
		testcase_3,
		testcase_4,
		testcase_5,
		testcase_6,
		testcase_7,
		testcase_8,
		testcase_9,
		testcase_10,
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
