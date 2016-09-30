package collections

type Entry struct {
	key   interface{}
	value interface{}
}

func NewEntry(key interface{}, value interface{}) *Entry {
	return &Entry{key, value}
}

func (e *Entry) getKey() interface{} {
	return e.key
}

func (e *Entry) getValue() interface{} {
	return e.value
}
