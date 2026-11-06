package piscine

// Reduce folds s left-to-right starting from its first rune.
func Reduce(s string, f func(rune, rune) rune) rune {
	runes := []rune(s)
	if len(runes) == 0 {
		return 0
	}
	acc := runes[0]
	for _, c := range runes[1:] {
		acc = f(acc, c)
	}
	return acc
}
