package allowedList

import "github.com/deckarep/golang-set"

type Whitelist struct {
	list mapset.Set
}

func NewWhitelist() *Whitelist {
	return &Whitelist{mapset.NewSet()}
}

func (b *Whitelist) Add(item string) bool {
	if item == "" {
		return false
	}

	return b.list.Add(item)
}

func (b *Whitelist) Remove(item string) bool {
	if !b.Has(item) {
		return false
	}

	b.list.Remove(item)
	return true
}

func (b *Whitelist) Has(item string) bool {
	return b.list.Contains(item)
}

func (b *Whitelist) Allowed(item string) bool {
	if b.list.Cardinality() == 0 {
		return true
	}

	return b.list.Contains(item)
}
