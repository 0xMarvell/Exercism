package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	//panic("Please implement the Keep function")
	var collection Ints
	for _, v := range i {
		if filter(v) {
			collection = append(collection, v)
		}
	}
	return collection
}

func (i Ints) Discard(filter func(int) bool) Ints {
	//panic("Please implement the Discard function")
	var collection Ints
	for _, v := range i {
		if !filter(v) {
			collection = append(collection, v)
		}
	}
	return collection
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	panic("Please implement the Keep function")
}

func (s Strings) Keep(filter func(string) bool) Strings {
	//panic("Please implement the Keep function")
	var collection Strings
	for _, v := range s {
		if filter(v) {
			collection = append(collection, v)
		}
	}
	return collection
}
