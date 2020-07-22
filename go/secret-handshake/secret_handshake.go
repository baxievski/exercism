// Package secret implements a solution for the exercism secret-handshake challenge
package secret

// Handshake converts a number to a series of strings
func Handshake(inp uint) []string {
	result := []string{}
	codes := map[uint]string{
		1: "wink",
		2: "double blink",
		4: "close your eyes",
		8: "jump",
	}
	for _, i := range []uint{1, 2, 4, 8} {
		if i&inp == i {
			result = append(result, codes[i])
		}
	}
	if 16&inp == 16 {
		result = reverse(result)
	}
	return result
}

func reverse(x []string) []string {
	l := len(x)
	result := make([]string, l)
	for i, v := range x {
		result[l-1-i] = v
	}
	return result
}
