package comparators

type UInt16Comparator struct {
}

func NewUInt16Comparator() *UInt16Comparator {
	return &UInt16Comparator{}
}

func (comparator *UInt16Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(uint16)
	bAsserted := b.(uint16)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
