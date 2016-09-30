package comparators

type Int16Comparator struct {
}

func NewInt16Comparator() *Int16Comparator {
	return &Int16Comparator{}
}

func (comparator *Int16Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(int16)
	bAsserted := b.(int16)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
