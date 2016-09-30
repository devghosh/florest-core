package comparators

type Int32Comparator struct {
}

func NewInt32Comparator() *Int32Comparator {
	return &Int32Comparator{}
}

func (comparator *Int32Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(int32)
	bAsserted := b.(int32)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
