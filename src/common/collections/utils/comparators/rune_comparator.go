package comparators

//runeComparator for comparing the rune values
type RuneComparator struct {
}

//Returns the new rune comparator
func NewRuneComparator() *RuneComparator {
	return &RuneComparator{}
}

// Compares two rune values and returns
// 0 if a = b
// -1 if a < b
// 1 if a > b
func (comparator *RuneComparator) Compare(a, b interface{}) int {
	aAsserted := a.(rune)
	bAsserted := b.(rune)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
