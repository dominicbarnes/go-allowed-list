package allowedList

import "github.com/deckarep/golang-set"

// Whitelist represents a strategy where only the listed items should be allowed.
type Whitelist struct {
	setlist
}

// NewWhitelist creates a new instance.
func NewWhitelist() *Whitelist {
	return &Whitelist{setlist{mapset.NewSet()}}
}

// Allowed tells us whether the input item should be allowed under this strategy.
func (white *Whitelist) Allowed(item string) bool {
	if white.Size() == 0 {
		return true
	}

	return white.Has(item)
}
