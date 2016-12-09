package allowedList

import "github.com/deckarep/golang-set"

type setlist struct {
	list mapset.Set
}

// Add includes the given item in the list. Empty strings are ignored and a bool
// is returned indicating whether or not the item was added.
func (set *setlist) Add(item string) bool {
	if item == "" {
		return false
	}

	return set.list.Add(item)
}

// Remove removes the given item from the list if it is present. Returns a bool
// indicating whether or not it removed the item.
func (set *setlist) Remove(item string) bool {
	if !set.Has(item) {
		return false
	}

	set.list.Remove(item)
	return true
}

// Has tells us whether the given item is included in the list.
func (set *setlist) Has(item string) bool {
	return set.list.Contains(item)
}

// Size tells us how many items are in the list.
func (set *setlist) Size() int {
	return set.list.Cardinality()
}
