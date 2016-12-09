package allowedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// assert that it implements the AllowedList interface.
func TestBlacklistImplements(t *testing.T) {
	assert.Implements(t, (*AllowedList)(nil), new(Blacklist))
}

// assert that empty strings are not added to the list.
func TestBlacklistAddEmptyString(t *testing.T) {
	b := NewBlacklist()
	b.Add("")
	assert.False(t, b.Has(""))
}

// assert that non-empty strings are added to the list.
func TestBlacklistAddSimple(t *testing.T) {
	b := NewBlacklist()
	b.Add("a")
	assert.True(t, b.Has("a"))
}

// assert that removing a non-existant item returns false.
func TestBlacklistRemoveFalse(t *testing.T) {
	b := NewBlacklist()
	assert.False(t, b.Remove("a"))
}

// assert that removing an item returns true.
func TestBlacklistRemoveTrue(t *testing.T) {
	b := NewBlacklist()
	b.Add("a")
	assert.True(t, b.Remove("a"))
}

// assert that items not added to the list are correctly identified.
func TestBlacklistHasFalse(t *testing.T) {
	b := NewBlacklist()
	assert.False(t, b.Has("a"))
}

// assert that items added to the list are correctly identified.
func TestBlacklistHasTrue(t *testing.T) {
	b := NewBlacklist()
	b.Add("a")
	assert.True(t, b.Has("a"))
}

// assert that an empty blacklist means anything is allowed.
func TestBlacklistAllowedEmpty(t *testing.T) {
	b := NewBlacklist()
	assert.True(t, b.Allowed("a"))
}

// assert that only items not in the blacklist are allowed.
func TestBlacklistAllowed(t *testing.T) {
	b := NewBlacklist()
	b.Add("a")
	assert.False(t, b.Allowed("a"))
	assert.True(t, b.Allowed("b"))
}
