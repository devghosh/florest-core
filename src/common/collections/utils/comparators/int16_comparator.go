package comparators

//int16Comparator for comparing the int16 values
type Int16Comparator struct {
}

//Returns the new int16 comparator
func NewInt16Comparator() *Int16Comparator {
	return &Int16Comparator{}
}

// Compares two int16 values and returns
// 0 if a = b
// -1 if a < b
// 1 if a > b
func (comparator *Int16Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(int16)
	bAsserted := b.(int16)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
