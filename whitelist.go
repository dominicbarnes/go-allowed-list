package allowedList

import "github.com/deckarep/golang-set"

var _ AllowedList = new(Whitelist)

// Whitelist represents a strategy where only the listed items should be allowed.
type Whitelist struct {
	setlist
	strict bool
}

// NewWhitelist creates a new instance.
func NewWhitelist(items ...string) *Whitelist {
	w := Whitelist{setlist: setlist{mapset.NewSet()}}
	w.Add(items...)
	return &w
}

// Strict enables "strict mode" for this whitelist. By default, an empty list
// will allow everything, but strict mode flips this and disallows everything.
func (white *Whitelist) Strict() {
	white.strict = true
}

// Allowed tells us whether the input item should be allowed under this strategy.
func (white *Whitelist) Allowed(item string) bool {
	if white.Size() == 0 {
		if white.strict {
			// in strict mode, we disallow everything when the list is empty
			return false
		}

		// by default, we allow everything when the list is empty
		return true
	}

	return white.Has(item)
}
