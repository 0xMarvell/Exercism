// Package strain provides basic separation functions.
package strain

type Ints []int
type Lists [][]int
type Strings []string

// Keep returns a new collection containing
// elements where the predicate is true.
func (i Ints) Keep(filter func(int) bool) Ints {
	var collection Ints

	for _, v := range i {
		if filter(v) {
			collection = append(collection, v)
		}
	}
	return collection
}

// Discard returns a new collection containing
// elements where the predicate is false.
func (i Ints) Discard(filter func(int) bool) Ints {
	var collection Ints

	for _, v := range i {
		if !filter(v) {
			collection = append(collection, v)
		}
	}
	return collection
}

// Keep returns a new collection containing
// elements where the predicate is true.
func (l Lists) Keep(filter func([]int) bool) Lists {
	var collection Lists

	for _, v := range l {
		if filter(v) {
			collection = append(collection, v)
		}
	}
	return collection
}

// Keep returns a new collection containing
// elements where the predicate is true.
func (s Strings) Keep(filter func(string) bool) Strings {
	var collection Strings

	for _, v := range s {
		if filter(v) {
			collection = append(collection, v)
		}
	}
	return collection
}
