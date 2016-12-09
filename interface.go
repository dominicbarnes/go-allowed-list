package allowedList

// AllowedList is the shared API for both Blacklist and Whitelist.
type AllowedList interface {
	Add(...string)
	Remove(...string)
	Has(string) bool
	Allowed(string) bool
}
