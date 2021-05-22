package util

// Set is a very simple which using the builtin map type and it is *not* threadsafe.
type Set struct {
	items map[interface{}]struct{}
}

func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		set.items[item] = struct{}{}
	}
}

func (set *Set) Remove(items ...interface{}) {
	for _, item := range items {
		delete(set.items, item)
	}
}

func (set *Set) Exists(item interface{}) bool {
	_, ok := set.items[item]

	return ok
}

func (set *Set) Len() int64 {
	size := int64(len(set.items))

	return size
}

func (set *Set) Clear() {
	set.items = map[interface{}]struct{}{}
}

func (set *Set) Keys() []interface{} {
	keys := make([]interface{}, 0, len(set.items))
	for item := range set.items {
		keys = append(keys, item)
	}
	return keys
}

func (set *Set) Difference(other *Set) *Set {
	difference := NewSet()
	for elem := range set.items {
		if !other.Exists(elem) {
			difference.Add(elem)
		}
	}
	return difference
}

func (set *Set) Union(other *Set) *Set {
	union := NewSet()
	for elem := range set.items {
		union.Add(elem)
	}
	for elem := range other.items {
		union.Add(elem)
	}
	return union
}

func (set *Set) Intersect(other *Set) *Set {
	intersection := NewSet()
	for elem := range set.items {
		if other.Exists(elem) {
			intersection.Add(elem)
		}
	}
	return intersection
}

func NewSet(items ...interface{}) *Set {
	set := &Set{}
	set.items = make(map[interface{}]struct{})
	for _, item := range items {
		set.items[item] = struct{}{}
	}

	return set
}
