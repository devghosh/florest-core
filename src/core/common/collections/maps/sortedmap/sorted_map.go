package sortedmap

type SortedMap interface {
	FirstKey() interface{}
    LastKey() interface{}
	Map, Iterable, Comparable
}