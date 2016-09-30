package comparators

type ByteComparator struct {
}

func NewByteComparator() *ByteComparator {
	return &ByteComparator{}
}

func (comparator *ByteComparator) Compare(a, b interface{}) int {
	aAsserted := a.(byte)
	bAsserted := b.(byte)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
