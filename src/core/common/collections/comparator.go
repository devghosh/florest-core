package collections

type Comparator interface {
	Compare(a, b interface{}) int
}
