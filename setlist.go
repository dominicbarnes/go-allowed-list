package allowedList

import "github.com/deckarep/golang-set"

type setlist struct {
	list mapset.Set
}

// Add includes the given item in the list. Empty strings are ignored.
func (set *setlist) Add(items ...string) {
	for _, item := range items {
		if item != "" {
			set.list.Add(item)
		}
	}
}

// Remove removes the given item from the list if it is present.
func (set *setlist) Remove(items ...string) {
	for _, item := range items {
		set.list.Remove(item)
	}
}

// Has tells us whether the given item is included in the list.
func (set *setlist) Has(item string) bool {
	return set.list.Contains(item)
}

// Size tells us how many items are in the list.
func (set *setlist) Size() int {
	return set.list.Cardinality()
}
