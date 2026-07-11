package piscine

import "fmt"

func SafeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			result = 0
			err = fmt.Errorf("cannot divide by zero")
		}
	}()
	result = a / b
	return
}

func WithCleanup(fn func(), cleanup func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
		cleanup()
	}()
	fn()
	return nil
}

func DeferOrder() (log []string) {
	defer func() { log = append(log, "first") }()
	defer func() { log = append(log, "second") }()
	defer func() { log = append(log, "third") }()
	return
}
