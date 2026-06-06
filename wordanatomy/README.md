# Word Anatomy

## Objective

Create a Go program that analyzes words and identifies known **prefixes** and **suffixes**.

The challenge is split into two files:

```text
main.go
wordanatomy.go
```

Your task is to implement the solution inside `wordanatomy.go`.

The correction program will use `main.go` to validate your implementation.

---

# Project Structure

## wordanatomy.go

This file must contain:

```go
package piscine
```

and the function:

```go
func WordAnatomy(word string) (string, string, string)
```

The function must return:

```text
(prefix, root, suffix)
```

---

## main.go

The testing file is provided by the checker.

You must **NOT modify it**.

The file contains all validation cases used during evaluation.

No additional files are required.

---

# Allowed Packages

None.

```text
No imports are allowed.
```

The entire solution must be implemented using only:

* string indexing
* loops
* conditionals
* functions

---

# Prefixes

The function must recognize the following prefixes:

```text
un
re
pre
mis
dis
over
under
anti
inter
sub
```

---

# Suffixes

The function must recognize the following suffixes:

```text
ing
ed
er
est
ly
ness
ment
tion
able
ful
```

---

# Expected Function

```go
func WordAnatomy(word string) (string, string, string)
```

### Return Values

| Position | Meaning         |
| -------- | --------------- |
| 1        | detected prefix |
| 2        | root word       |
| 3        | detected suffix |

---

# Detection Rules

## Prefix Only

Input:

```text
unhappy
```

Return:

```go
"un", "happy", ""
```

---

## Suffix Only

Input:

```text
runner
```

Return:

```go
"", "runn", "er"
```

---

## Prefix and Suffix

Input:

```text
redoing
```

Return:

```go
"re", "do", "ing"
```

---

## No Match

Input:

```text
planet
```

Return:

```go
"", "planet", ""
```

---

# Matching Rules

### Prefix

A prefix must appear at the beginning of the word.

Example:

```text
misunderstand
```

returns:

```go
"mis", "understand", ""
```

---

### Suffix

A suffix must appear at the end of the word.

Example:

```text
careful
```

returns:

```go
"", "care", "ful"
```

---

### Prefix + Suffix

Both may exist simultaneously.

Example:

```text
misunderstanding
```

returns:

```go
"mis", "understand", "ing"
```

---

# Longest Match Rule

If multiple prefixes could match, use the longest one.

Example:

```text
underpaid
```

returns:

```go
"under", "paid", ""
```

not

```go
"un", "derpaid", ""
```

---

Likewise for suffixes.

Example:

```text
greatest
```

returns:

```go
"", "great", "est"
```

not

```go
"", "greatest", ""
```

---

# main.go (Checker)

The correction system will use a file similar to the following:

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	tests := []string{
		"unhappy",
		"runner",
		"redoing",
		"kindness",
		"planet",
		"careful",
		"misunderstanding",
		"greatest",
		"underpaid",
	}

	for _, word := range tests {
		prefix, root, suffix := piscine.WordAnatomy(word)

		fmt.Printf(
			"%s -> Prefix:%q Root:%q Suffix:%q\n",
			word,
			prefix,
			root,
			suffix,
		)
	}
}
```

---

# Expected Output

```text
unhappy -> Prefix:"un" Root:"happy" Suffix:""
runner -> Prefix:"" Root:"runn" Suffix:"er"
redoing -> Prefix:"re" Root:"do" Suffix:"ing"
kindness -> Prefix:"" Root:"kind" Suffix:"ness"
planet -> Prefix:"" Root:"planet" Suffix:""
careful -> Prefix:"" Root:"care" Suffix:"ful"
misunderstanding -> Prefix:"mis" Root:"understand" Suffix:"ing"
greatest -> Prefix:"" Root:"great" Suffix:"est"
underpaid -> Prefix:"under" Root:"paid" Suffix:""
```

---

# Examples

### Example 1

Input:

```go
WordAnatomy("unhappy")
```

Output:

```go
"un", "happy", ""
```

---

### Example 2

Input:

```go
WordAnatomy("redoing")
```

Output:

```go
"re", "do", "ing"
```

---

### Example 3

Input:

```go
WordAnatomy("planet")
```

Output:

```go
"", "planet", ""
```

---

# Bonus

### Level 1

Support multiple prefixes.

```text
antidisestablishment
```

returns:

```go
"anti-dis", "establishment", ""
```

---

### Level 2

Support multiple suffixes.

```text
carefulness
```

returns:

```go
"", "care", "ful-ness"
```

---

### Level 3

Return the anatomy visually.

```text
redoing
```

becomes:

```text
[re][do][ing]
```

---

A good linguist can identify words.

A great programmer can dissect them.