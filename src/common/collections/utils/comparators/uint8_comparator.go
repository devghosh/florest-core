package comparators

type UInt8Comparator struct {
}

func NewUInt8Comparator() *UInt8Comparator {
	return &UInt8Comparator{}
}

func (comparator *UInt8Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(uint8)
	bAsserted := b.(uint8)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
