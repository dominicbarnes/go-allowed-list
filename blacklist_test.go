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

// assert that items can be removed.
func TestBlacklistRemove(t *testing.T) {
	b := NewBlacklist()
	b.Add("a")
	b.Remove("a")
	assert.False(t, b.Has("a"))
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

// assert that the size is accurately reported.
func TestBlacklistSizeEmpty(t *testing.T) {
	b := NewBlacklist()
	assert.Zero(t, b.Size())
	b.Add("a")
	assert.Equal(t, 1, b.Size())
}
