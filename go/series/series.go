package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	result := make([]string, 0, len(s)-n+1)
	for i := 0; i <= len(s)-n; i++ {
		if f, ok := First(n, s[i:]); ok {
			result = append(result, f)
		}
	}
	return result
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	f, _ := First(n, s)
	return f
}

// First returns the first substring of s with length n.
func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
