package allowedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// assert that it implements the AllowedList interface.
func TestWhitelistImplements(t *testing.T) {
	assert.Implements(t, (*AllowedList)(nil), new(Whitelist))
}

// assert that empty strings are not added to the list.
func TestWhitelistAddEmptyString(t *testing.T) {
	b := NewWhitelist()
	b.Add("")
	assert.False(t, b.Has(""))
}

// assert that non-empty strings are added to the list.
func TestWhitelistAddSimple(t *testing.T) {
	b := NewWhitelist()
	b.Add("a")
	assert.True(t, b.Has("a"))
}

// assert that removing a non-existant item returns false.
func TestWhitelistRemoveFalse(t *testing.T) {
	b := NewWhitelist()
	assert.False(t, b.Remove("a"))
}

// assert that removing an item returns true.
func TestWhitelistRemoveTrue(t *testing.T) {
	b := NewWhitelist()
	b.Add("a")
	assert.True(t, b.Remove("a"))
}

// assert that items not added to the list are correctly identified.
func TestWhitelistHasFalse(t *testing.T) {
	b := NewWhitelist()
	assert.False(t, b.Has("a"))
}

// assert that items added to the list are correctly identified.
func TestWhitelistHasTrue(t *testing.T) {
	b := NewWhitelist()
	b.Add("a")
	assert.True(t, b.Has("a"))
}

// assert that an empty whitelist means anything is allowed.
func TestWhitelistAllowedEmpty(t *testing.T) {
	b := NewWhitelist()
	assert.True(t, b.Allowed("a"))
}

// assert that only items in the whitelist are allowed.
func TestWhitelistAllowed(t *testing.T) {
	b := NewWhitelist()
	b.Add("a")
	assert.True(t, b.Allowed("a"))
	assert.False(t, b.Allowed("b"))
}
