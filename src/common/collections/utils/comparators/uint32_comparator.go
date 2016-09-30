package comparators

type UInt32Comparator struct {
}

func NewUInt32Comparator() *UInt32Comparator {
	return &UInt32Comparator{}
}

func (comparator *UInt32Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(uint32)
	bAsserted := b.(uint32)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
