package comparators

type Int64Comparator struct {
}

func NewInt64Comparator() *Int64Comparator {
	return &Int64Comparator{}
}

func (comparator *Int64Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(int64)
	bAsserted := b.(int64)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
