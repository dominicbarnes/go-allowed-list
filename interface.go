// Package allowedList is a library for creating allowed lists using either the
// blacklist (disallow items) and whitelist (only allow items) strategies.
package allowedList

// AllowedList is the shared API for both Blacklist and Whitelist.
type AllowedList interface {
	Add(...string)
	Remove(...string)
	Has(string) bool
	Allowed(string) bool
}
