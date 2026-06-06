## rpncalc

### Instructions

Write a program that takes a `string` which contains an equation written in `Reverse Polish Notation` (RPN) as its first argument, that evaluates the equation, and that prints the result on the standard output followed by a newline (`'\n'`).

`Reverse Polish Notation` is a mathematical notation in which every operator follows all of its operands. In RPN, every operator encountered evaluates the previous 2 operands, and the result of this operation then becomes the first of the two operands for the subsequent operator. Operands and operators must be spaced by at least one space.

The following operators must be implemented : `+`, `-`, `*`, `/`, and `%`.

If the `string` is not valid or if there is not exactly one argument, `Error` must be printed on the standard output followed by a newline.
If the `string` has extra spaces it is still considered valid.

All the given operands must fit in a `int`.

### Usage

```console
$ go run . "1 2 * 3 * 4 +" | cat -e
10$
$ go run . "1 2 3 4 +" | cat -e
Error$
$ go run . | cat -e
Error$
$ go run . "     1      3 * 2 -" | cat -e
1
$ go run . "     1      3 * ksd 2 -" | cat -e
Error$
$
```
