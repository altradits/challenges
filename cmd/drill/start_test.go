package main

import (
	"strings"
	"testing"
)

func TestExtractTaskTemplate(t *testing.T) {
	readme := `# 06-functions — Greet

## Task

Write Greet.

## Input / Output

One string in, one out.

## Examples

Greet("Go") → "Hello, Go!"

## Constraints

Exported name.

## Time limit

**15 minutes** (round 1).

## Blind rule

No peeking.

## How to check

drill check 6
`
	got := extractTask(readme)
	for _, want := range []string{"## Task", "Write Greet.", "## Examples", "Greet(\"Go\")", "## Constraints", "## Time limit", "**15 minutes**"} {
		if !strings.Contains(got, want) {
			t.Errorf("extractTask missing %q\n--- got ---\n%s", want, got)
		}
	}
	for _, banned := range []string{"Blind rule", "No peeking", "How to check"} {
		if strings.Contains(got, banned) {
			t.Errorf("extractTask leaked %q", banned)
		}
	}
}

func TestExtractTaskLegacy(t *testing.T) {
	readme := `## gcd

### Instructions

Write Gcd(a, b uint) uint.

### Expected function

` + "```go\nfunc Gcd(a, b uint) uint\n```" + `

### Usage

Here is a possible program to test your function.
`
	got := extractTask(readme)
	for _, want := range []string{"### Instructions", "Write Gcd", "### Expected function", "func Gcd"} {
		if !strings.Contains(got, want) {
			t.Errorf("extractTask missing %q\n--- got ---\n%s", want, got)
		}
	}
}

func TestExtractTaskEmpty(t *testing.T) {
	if got := extractTask("# hi\n\nno sections\n"); !strings.Contains(got, "no task section") {
		t.Fatalf("want placeholder, got %q", got)
	}
}
