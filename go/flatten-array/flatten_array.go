// Package flatten implements a solution for the flatten array exercism problem
package flatten

// Flatten gives a flat array of non nil values from an arbitrarily deep nested array
func Flatten(nested interface{}) []interface{} {
	flat := []interface{}{}

	for _, val := range nested.([]interface{}) {
		switch val.(type) {
		case []interface{}:
			flat = append(flat, Flatten(val)...)
		case nil:
		default:
			flat = append(flat, val)
		}
	}

	return flat
}
