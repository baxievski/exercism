package strain

// Ints is a collection of integers
type Ints []int

// Lists is a collection of lists of integers
type Lists [][]int

// Strings is a collection of strings
type Strings []string

// Keep returns a collection containing the elements for which the predicate is true
func (l Ints) Keep(fn func(int) bool) Ints {
	if l == nil {
		return nil
	}
	result := Ints{}
	for _, e := range l {
		if !fn(e) {
			continue
		}
		result = append(result, e)
	}
	return result
}

// Discard returns a collection containing the elements for which the predicate is false
func (l Ints) Discard(fn func(int) bool) Ints {
	if l == nil {
		return nil
	}
	result := Ints{}
	for _, e := range l {
		if fn(e) {
			continue
		}
		result = append(result, e)
	}
	return result
}

// Keep returns a collection containing the elements for which the predicate is true
func (l Lists) Keep(fn func([]int) bool) Lists {
	result := Lists{}
	for _, e := range l {
		if !fn(e) {
			continue
		}
		result = append(result, e)
	}
	return result
}

// Discard returns a collection containing the elements for which the predicate is false
func (l Lists) Discard(fn func([]int) bool) Lists {
	result := Lists{}
	for _, e := range l {
		if fn(e) {
			continue
		}
		result = append(result, e)
	}
	return result
}

// Keep returns a collection containing the elements for which the predicate is true
func (l Strings) Keep(fn func(string) bool) Strings {
	result := Strings{}
	for _, e := range l {
		if !fn(e) {
			continue
		}
		result = append(result, e)
	}
	return result
}

// Discard returns a collection containing the elements for which the predicate is false
func (l Strings) Discard(fn func(string) bool) Strings {
	result := Strings{}
	for _, e := range l {
		if fn(e) {
			continue
		}
		result = append(result, e)
	}
	return result
}
