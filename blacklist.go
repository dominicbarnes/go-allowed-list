package allowedList

import "github.com/deckarep/golang-set"

// Blacklist represents a strategy where the listed items should be rejected.
type Blacklist struct {
	setlist
}

// NewBlacklist creates a new instance.
func NewBlacklist() *Blacklist {
	return &Blacklist{setlist{mapset.NewSet()}}
}

// Allowed tells us whether the input item should be allowed under this strategy.
func (black *Blacklist) Allowed(item string) bool {
	if black.Size() == 0 {
		return true
	}

	return !black.Has(item)
}
