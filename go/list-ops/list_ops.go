// Package listops implements a solution for the exercism list-ops challenge
package listops

// IntList represent list of integers
type IntList []int

type predFunc func(int) bool

type unaryFunc func(int) int

type binFunc func(x, y int) int

// Foldr folds each item of the list into the accumulator from the right using binFunc(item, acc)
func (l IntList) Foldr(fn binFunc, acc int) int {
	result := acc
	for _, v := range l.Reverse() {
		result = fn(v, result)
	}
	return result
}

// Foldl folds each item of the list into the accumulator from the left using binFunc(acc, item)
func (l IntList) Foldl(fn binFunc, acc int) int {
	result := acc
	for _, v := range l {
		result = fn(result, v)
	}
	return result
}

// Map gives a list of the results of applying unaryFunc to each of the elements of the list
func (l IntList) Map(fn unaryFunc) IntList {
	result := make(IntList, l.Length())
	for i, v := range l {
		result[i] = fn(v)
	}
	return result
}

// Filter gives a list of the elements of the list for which predFunc(element) is true
func (l IntList) Filter(fn predFunc) IntList {
	len := 0
	for _, v := range l {
		if fn(v) {
			len++
		}
	}
	result := make(IntList, len)
	i := 0
	for _, v := range l {
		if fn(v) {
			result[i] = v
			i++
		}
	}
	return result
}

// Reverse gives a list of the same item as the original list, but in reverse order
func (l IntList) Reverse() IntList {
	length := l.Length()
	result := make(IntList, length)
	for i, v := range l {
		result[length-i-1] = v
	}
	return result
}

// Length gives the lenght of the list
func (l IntList) Length() int {
	result := 0
	for range l {
		result++
	}
	return result
}

// Append adds all the elements of the other list at the end of l
func (l IntList) Append(other IntList) IntList {
	len1 := l.Length()
	result := make(IntList, len1+other.Length())
	for i, v1 := range l {
		result[i] = v1
	}
	for j, v2 := range other {
		result[len1+j] = v2
	}
	return result
}

// Concat flattens all the lists into one
func (l IntList) Concat(others []IntList) IntList {
	result := l
	for _, other := range others {
		result = result.Append(other)
	}
	return result
}
