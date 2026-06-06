# Word Anatomy 2

## Objective

Write a function `Wordanatomy()` that takes a word, a prefix array, and a suffix array. The function should return the first matching prefix or suffix for the word based on the provided arrays.

---

## Expected Function

```go
func Wordanatomy(word string, pref []string, suff []string) string
```

### Parameters

| Parameter | Type     | Description                          |
|-----------|----------|--------------------------------------|
| word      | string   | The word to analyze                  |
| pref      | []string | Array of prefixes to check           |
| suff      | []string | Array of suffixes to check           |

### Return Value

Returns a formatted string:

```text
prefix: "prefix", suffix: "suffix"
```

---

## Rules

### Both Prefix and Suffix

If the word has both a matching prefix and suffix:

```text
prefix: "un", suffix: "able"
```

---

### Only Prefix

If the word has only a matching prefix:

```text
prefix: "un", suffix: ""
```

---

### Only Suffix

If the word has only a matching suffix:

```text
prefix: "", suffix: "able"
```

---

### No Match

If the word has no matching prefix or suffix:

```text
prefix: "", suffix: ""
```

---

## Matching Rules

### Prefix

A prefix must appear at the **beginning** of the word.

Example with prefixes `["an", "en", "un"]`:

```text
understandable
```

The first matching prefix is `"un"` (not `"an"` or `"en"`).

---

### Suffix

A suffix must appear at the **end** of the word.

Example with suffixes `["ible", "able"]`:

```text
understandable
```

The first matching suffix is `"able"` (not `"ible"`).

---

## First Match Rule

Use the **first** matching prefix/suffix from the provided arrays (not the longest).

Example:

```text
prefixes: ["an", "en", "un"]
word: "understandable"
```

Returns:

```text
prefix: "un", suffix: "able"
```

---

## Example Usage

```go
package main

import (
	"fmt"
)

func Wordanatomy(word string, pref []string, suff []string) string {
	// Your implementation here
}

func main() {
	prefixes := []string{"an", "en", "un"}
	suffixes := []string{"ible", "able"}

	// Test Case 1: Both Prefix and Suffix exist
	fmt.Println(Wordanatomy("understandable", prefixes, suffixes))
	// Output: prefix: "un", suffix: "able"

	// Test Case 2: Only Prefix exists
	fmt.Println(Wordanatomy("understand", prefixes, suffixes))
	// Output: prefix: "un", suffix: ""

	// Test Case 3: Neither exists
	fmt.Println(Wordanatomy("apple", prefixes, suffixes))
	// Output: prefix: "", suffix: ""
}
```

---

## Expected Output

```text
prefix: "un", suffix: "able"
prefix: "un", suffix: ""
prefix: "", suffix: ""