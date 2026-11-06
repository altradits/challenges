package piscine

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return nil
	}

	start := nbrs[0]
	if start < 0 {
		start = len(a) + start
	}
	if start < 0 {
		start = 0
	}
	if start > len(a) {
		return nil
	}

	if len(nbrs) == 1 {
		return a[start:]
	}

	end := nbrs[1]
	if end < 0 {
		end = len(a) + end
	}
	if end < 0 {
		end = 0
	}
	if end > len(a) {
		end = len(a)
	}
	if end < start {
		return nil
	}

	return a[start:end]
}
