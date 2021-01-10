package allergies

import "sort"

var allergenValues = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

// Allergies gives a list of all the allergens besed on the allergy score
func Allergies(score uint) []string {
	result := []string{}
	for _, value := range sortedKeys(allergenValues) {
		if (score & value) == 0 {
			continue
		}
		result = append(result, allergenValues[value])
	}
	return result
}

// AllergicTo tells if a person is allergic to a substance based on their allergy score
func AllergicTo(score uint, substance string) bool {
	for _, allergen := range Allergies(score) {
		if substance == allergen {
			return true
		}
	}
	return false
}

type allergenScores []uint

func (a allergenScores) Len() int {
	return len(a)
}

func (a allergenScores) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a allergenScores) Less(i, j int) bool {
	return a[i] < a[j]
}

func sortedKeys(m map[uint]string) []uint {
	keys := make(allergenScores, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(keys)
	return keys
}
