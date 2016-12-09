package allowedList

// AllowedList is the shared API for both Blacklist and Whitelist.
type AllowedList interface {
	Add(string) bool
	Remove(string) bool
	Has(string) bool
	Allowed(string) bool
}
