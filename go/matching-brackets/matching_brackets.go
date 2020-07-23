// Package brackets implements a solution for the exercism matching-brackets challenge
package brackets

// Bracket tells if the brackets are matched
func Bracket(brackets string) bool {
	stack := []rune{}
	openBrackets := map[rune]bool{
		'{': true,
		'[': true,
		'(': true,
	}
	cloosedToOpen := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	for _, b := range clean(brackets) {
		if openBrackets[b] {
			stack = append(stack, b)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		lastBrace := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		openingBrace := cloosedToOpen[b]
		if lastBrace != openingBrace {
			return false
		}
	}
	return len(stack) == 0
}

func clean(braces string) string {
	allBraces := map[rune]bool{
		'{': true,
		'[': true,
		'(': true,
		'}': true,
		']': true,
		')': true,
	}
	cleaned := []rune{}
	for _, v := range braces {
		if allBraces[v] {
			cleaned = append(cleaned, v)
		}
	}
	return string(cleaned)
}
