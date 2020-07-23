// Package brackets implements a solution for the exercism matching-brackets challenge
package brackets

// Bracket tells if the brackets are matched
func Bracket(brackets string) bool {
	stack := []rune{}
	cloosedToOpen := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	for _, b := range brackets {
		switch b {
		case '{', '[', '(':
			stack = append(stack, b)
		case '}', ']', ')':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != cloosedToOpen[b] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
