package piscine

func Map[T, U any](s []T, fn func(T) U) []U {
	out := make([]U, len(s))
	for i, v := range s {
		out[i] = fn(v)
	}
	return out
}

func Filter[T any](s []T, fn func(T) bool) []T {
	var out []T
	for _, v := range s {
		if fn(v) {
			out = append(out, v)
		}
	}
	return out
}

func Reduce[T, U any](s []T, initial U, fn func(U, T) U) U {
	acc := initial
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Contains[T comparable](s []T, target T) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}
