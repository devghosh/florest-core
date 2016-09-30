package comparators

type UInt64Comparator struct {
}

func NewUInt64Comparator() *UInt64Comparator {
	return &UInt64Comparator{}
}

func (comparator *UInt64Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(uint64)
	bAsserted := b.(uint64)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
