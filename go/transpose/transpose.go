// Package transpose implements a solution for the exercism transpose challenge
package transpose

// Transpose gives the trannsposed output oof a text
func Transpose(tx []string) []string {
	input := make([][]rune, len(tx))
	longest := 0
	for i, s := range tx {
		input[i] = []rune(s)
		if (len(input[i])) > longest {
			longest = len(input[i])
		}
	}

	transposed := make([][]rune, longest)

	for _, runes := range pad(input) {
		for j := 0; j < longest; j++ {
			if j < len(runes) {
				transposed[j] = append(transposed[j], runes[j])
			}
		}
	}

	result := make([]string, longest)
	for i, runes := range transposed {
		result[i] = string(runes)
	}

	return result
}

func pad(input [][]rune) [][]rune {
	lengths := make([]int, len(input))
	for i, runes := range input {
		lengths[i] = len(runes)
	}
	padded := make([][]rune, len(input))
	m := 0
	for i, runes := range input {
		m = max(lengths[i+1:])
		if lengths[i] >= m {
			padded[i] = runes
			continue
		}
		l := make([]rune, m)
		for j := range l {
			if j < lengths[i] {
				l[j] = runes[j]
				continue
			}
			l[j] = ' '
		}
		padded[i] = l
	}
	return padded
}

func max(l []int) int {
	m := 0
	for _, v := range l {
		if v > m {
			m = v
		}
	}
	return m
}
