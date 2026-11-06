// Command drill is the memory-first trainer for the Go Mastery challenges.
//
// It enforces the method: start a timed blind attempt, write the solution
// from memory, check it against the hidden tests, and repeat on a spaced
// schedule. Standard library only, offline, no accounts: progress lives in
// .drill/progress.json at the repo root (gitignored).
//
// Usage:
//
//	drill start NN     start the timer, show only the task
//	drill check NN     run the tests, record pass/fail and time
//	drill due          list what spaced repetition says to redo today
//	drill stats        streaks, times, weak topics
//	drill solution NN  show the reference solution (only after passing once)
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	root, err := findRoot()
	if err != nil {
		fatal(err)
	}
	var cmdErr error
	switch os.Args[1] {
	case "start":
		cmdErr = cmdStart(root, os.Args[2:])
	case "check":
		cmdErr = cmdCheck(root, os.Args[2:])
	case "due":
		cmdErr = cmdDue(root, os.Args[2:])
	case "stats":
		cmdErr = cmdStats(root, os.Args[2:])
	case "solution":
		cmdErr = cmdSolution(root, os.Args[2:])
	case "help", "-h", "--help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n\n", os.Args[1])
		usage()
		os.Exit(2)
	}
	if cmdErr != nil {
		fatal(cmdErr)
	}
}

func usage() {
	fmt.Println(`drill — code Go from memory.

Usage:
  drill start NN     start the timer, show only the task
  drill check NN     run the tests, record pass/fail and time
  drill due          list what spaced repetition says to redo today
  drill stats        streaks, times, weak topics
  drill solution NN  show the reference solution (only after passing once)

NN is the challenge number, with or without leading zeros (6 and 06 match).`)
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "drill:", err)
	os.Exit(1)
}
