package comparators

//float32Comparator for comparing the float32 values
type Float32Comparator struct {
}

//Returns the new float32 comparator
func NewFloat32Comparator() *Float32Comparator {
	return &Float32Comparator{}
}

// Compares two float32 values and returns
// 0 if a = b
// -1 if a < b
// 1 if a > b
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
