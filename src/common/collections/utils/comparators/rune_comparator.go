package comparators

type RuneComparator struct {
}

func NewRuneComparator() *RuneComparator {
	return &RuneComparator{}
}

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
