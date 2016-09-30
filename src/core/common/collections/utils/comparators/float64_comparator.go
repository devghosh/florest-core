package comparators

type Float64Comparator struct {
}

func NewFloat64Comparator() *Float64Comparator {
	return &Float64Comparator{}
}

func (comparator *Float64Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(float64)
	bAsserted := b.(float64)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
