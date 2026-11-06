package piscine

func Map[T, U any](s []T, fn func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = fn(v)
	}
	return result
}

func Filter[T any](s []T, keep func(T) bool) []T {
	result := []T{}
	for _, v := range s {
		if keep(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T, U any](s []T, initial U, fn func(U, T) U) U {
	acc := initial
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}

func Contains[T comparable](s []T, target T) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Min[T interface{ ~int | ~float64 | ~string }](s []T) T {
	m := s[0]
	for _, v := range s[1:] {
		if v < m {
			m = v
		}
	}
	return m
}
