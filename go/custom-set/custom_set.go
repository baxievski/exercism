package stringset

import (
	"fmt"
	"strings"
)

// Set represents a custom set type
type Set struct {
	elements map[string]bool
}

// New cretes an empty set
func New() Set {
	return Set{elements: map[string]bool{}}
}

// NewFromSlice creates a set with elements of a slice
func NewFromSlice(elements []string) Set {
	s := Set{elements: map[string]bool{}}
	for _, el := range elements {
		s.Add(el)
	}
	return s
}

// Add inserts an element into the set
func (s Set) Add(el string) {
	s.elements[el] = true
}

// String gives the string representation of the set
func (s Set) String() string {
	sb := strings.Builder{}
	sb.WriteString(`{`)
	i := -1
	lastIndex := len(s.elements) - 1
	for el := range s.elements {
		i++
		if i < lastIndex {
			sb.WriteString(fmt.Sprintf(`"%s", `, el))
			continue
		}
		sb.WriteString(fmt.Sprintf(`"%s"`, el))
	}
	sb.WriteString(`}`)
	return sb.String()
}

// IsEmpty tells if the set contains no elements
func (s Set) IsEmpty() bool {
	return len(s.elements) == 0
}

// Has tells if an element is contained in the set
func (s *Set) Has(el string) bool {
	if _, ok := s.elements[el]; ok {
		return true
	}
	return false
}

// Subset tellls if s1 is a subset of s2
func Subset(s1, s2 Set) bool {
	for el := range s1.elements {
		if !s2.Has(el) {
			return false
		}
	}
	return true
}

// Disjoint tells if the two sets have no elements in common
func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).IsEmpty()
}

// Equal tells if the two sets have the same elements
func Equal(s1, s2 Set) bool {
	if !Subset(s1, s2) {
		return false
	}
	if !Subset(s2, s1) {
		return false
	}
	return true
}

// Intersection gives a set of only the elements that are both in s1 and s2
func Intersection(s1, s2 Set) Set {
	result := New()
	for el := range s1.elements {
		if s2.Has(el) {
			result.Add(el)
		}
	}
	return result
}

// Difference gives a set of all the elements that are in s1, but not in s2
func Difference(s1, s2 Set) Set {
	result := New()
	for el := range s1.elements {
		if !s2.Has(el) {
			result.Add(el)
		}
	}
	return result
}

// Union gives a new set containing the distinct elements in eiter of the two sets
func Union(s1, s2 Set) Set {
	result := New()
	for _, s := range []Set{s1, s2} {
		for el := range s.elements {
			result.Add(el)
		}
	}
	return result
}
