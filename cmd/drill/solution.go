package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func cmdSolution(root string, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: drill solution NN")
	}
	cat, err := loadCatalog(root)
	if err != nil {
		return err
	}
	c, err := findChallenge(cat, args[0])
	if err != nil {
		return err
	}
	prog, err := loadProgress(root)
	if err != nil {
		return err
	}
	if prog.record(c.Num).Passes == 0 {
		return fmt.Errorf("solution for %s stays hidden until you pass it once (drill check %d)", c.Dir, c.Num)
	}
	path, err := solutionPath(root, c.Dir)
	if err != nil {
		return err
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Printf("=== reference solution: %s ===\n\n", path)
	fmt.Println(string(raw))
	return nil
}

// solutionPath finds the reference solution, preferring solution.go.
func solutionPath(root, dir string) (string, error) {
	for _, name := range []string{"solution.go", "main.go"} {
		p := filepath.Join(root, "solutions", dir, name)
		if st, err := os.Stat(p); err == nil && !st.IsDir() {
			return p, nil
		}
	}
	return "", fmt.Errorf("no reference solution stored for %s yet", dir)
}
