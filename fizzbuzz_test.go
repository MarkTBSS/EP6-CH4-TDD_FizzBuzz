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
	testcase_0 := Case{name: "FizzBuzz(1)", number: 1, want: "1"}
	tests := []Case{testcase_0}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			got := FizzBuzz(testcase.number)
			if got != testcase.want {
				t.Errorf("FizzBuzz(%d) = %s; want %s", testcase.number, got, testcase.want)
			}
		})
	}
}
