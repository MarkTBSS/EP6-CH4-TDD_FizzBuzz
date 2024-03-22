```
# FizzBuzz
  input  => output
    1    => "1"
    2    => "2"
    3    => "Fizz"
    4    => "4"
    5    => "Buzz"
    6    => "Fizz"
    7    => "7"
    8    => "8"
    9    => "Fizz"
    10   => "Buzz"
    11   => "11"
    12   => "Fizz"
    13   => "13"
    14   => "14"
    15   => "FizzBuzz"
    30   => "FizzBuzz"

check git commit to see step of TDD

go test -v .
go test -v -cover .
go test -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html
```