// Package sublist implements a solution for the exercism sublist challenge
package sublist

const equal = "equal"
const sublist = "sublist"
const superlist = "superlist"
const unequal = "unequal"

// Relation represents the relationship
type Relation string

// Sublist gives the sublist/superlist/equal/unequal relationship between two lists
func Sublist(a, b []int) Relation {
	if len(a) == len(b) {
		if isEqual(a, b) {
			return Relation(equal)
		}
		return Relation(unequal)
	}
	if len(a) < len(b) {
		if isSublist(a, b) {
			return Relation(sublist)
		}
		return Relation(unequal)
	}
	if isSublist(b, a) {
		return Relation(superlist)
	}
	return Relation(unequal)
}

func isEqual(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func isSublist(x, y []int) bool {
	if len(x) >= len(y) {
		return isEqual(x, y)
	}
	same := true
	for i := 0; i <= len(y)-len(x); i++ {
		same = true
		for j, v := range x {
			if v != y[i+j] {
				same = false
				break
			}
		}
		if same {
			return true
		}
	}
	return false
}
