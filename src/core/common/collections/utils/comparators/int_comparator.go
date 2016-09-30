package comparators

type IntComparator struct {
}

func NewIntComparator() *IntComparator {
	return &IntComparator{}
}

func (comparator *IntComparator) Compare(a, b interface{}) int {
	aAsserted := a.(int)
	bAsserted := b.(int)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
