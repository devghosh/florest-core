package comparators

type Float32Comparator struct {
}

func NewFloat32Comparator() *Float32Comparator {
	return &Float32Comparator{}
}

func (comparator *Float32Comparator) Compare(a, b interface{}) int {
	aAsserted := a.(float32)
	bAsserted := b.(float32)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}
