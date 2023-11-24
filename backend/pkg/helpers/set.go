package helpers

type StringSet map[string]struct{}

func MakeStringSet(items ...string) StringSet {
	set := make(StringSet)
	for _, item := range items {
		set[item] = struct{}{}
	}
	return set
}

func (s *StringSet) Contains(status string) bool {
	_, ok := (*s)[status]
	return ok
}
