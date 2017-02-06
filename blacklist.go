package allowedList

import "github.com/deckarep/golang-set"

var _ AllowedList = new(Blacklist)

// Blacklist represents a strategy where the listed items should be rejected.
type Blacklist struct {
	setlist
	strict bool
}

// NewBlacklist creates a new instance.
func NewBlacklist(items ...string) *Blacklist {
	b := Blacklist{setlist: setlist{mapset.NewSet()}}
	b.Add(items...)
	return &b
}

// Strict enables "strict mode" for this blacklist. By default, an empty list
// will allow everything, but strict mode flips this and disallows everything.
func (black *Blacklist) Strict() {
	black.strict = true
}

// Allowed tells us whether the input item should be allowed under this strategy.
func (black *Blacklist) Allowed(item string) bool {
	if black.Size() == 0 {
		if black.strict {
			// in strict mode, we disallow everything when the list is empty
			return false
		}

		// by default, we allow everything when the list is empty
		return true
	}

	return !black.Has(item)
}
