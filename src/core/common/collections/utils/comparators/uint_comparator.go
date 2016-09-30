package comparators

type UIntComparator struct {
}

func NewUIntComparator() *UIntComparator {
	return &UIntComparator{}
}

func (comparator *UIntComparator) Compare(a, b interface{}) int {
	aAsserted := a.(uint)
	bAsserted := b.(uint)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
