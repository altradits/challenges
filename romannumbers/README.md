## romannumbers

### Instructions

Write a program called `rn`. The objective is to convert a number, given as an argument, into a roman number and print it with its roman number calculation.

The program should have a limit of `4000`. In case of an invalid number, for example `"hello"` or `0` the program should print `ERROR: cannot convert to roman digit`.

### Usage

```console
$ go run . hello
ERROR: cannot convert to roman digit
$ go run . 123
C+X+X+I+I+I
CXXIII
$ go run . 999
(M-C)+(C-X)+(X-I)
CMXCIX
$ go run . 3999
M+M+M+(M-C)+(C-X)+(X-I)
MMMCMXCIX
$ go run . 4000
ERROR: cannot convert to roman digit
$
```
