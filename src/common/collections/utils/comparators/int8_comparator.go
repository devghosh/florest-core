package comparators

type Int8Comparator struct {
}

func NewInt8Comparator() *Int8Comparator {
	return &Int8Comparator{}
}

func (comparator *Int8Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(int8)
	bAsserted := b.(int8)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
