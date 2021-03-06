package allowedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// assert that it implements the AllowedList interface.
func TestWhitelistImplements(t *testing.T) {
	assert.Implements(t, (*AllowedList)(nil), new(Whitelist))
}

// assert that the items can be added on init.
func TestNewWhitelistItems(t *testing.T) {
	b := NewWhitelist("a", "b")
	assert.True(t, b.Has("a"))
	assert.True(t, b.Has("b"))
	assert.False(t, b.Has("c"))
}

// assert that empty strings are not added to the list.
func TestWhitelistAddEmptyString(t *testing.T) {
	w := NewWhitelist()
	w.Add("")
	assert.False(t, w.Has(""))
}

// assert that non-empty strings are added to the list.
func TestWhitelistAddSimple(t *testing.T) {
	w := NewWhitelist()
	w.Add("a")
	assert.True(t, w.Has("a"))
}

// assert that items can be removed.
func TestWhitelistRemove(t *testing.T) {
	w := NewWhitelist()
	w.Add("a")
	w.Remove("a")
	assert.False(t, w.Has("a"))
}

// assert that items not added to the list are correctly identified.
func TestWhitelistHasFalse(t *testing.T) {
	w := NewWhitelist()
	assert.False(t, w.Has("a"))
}

// assert that items added to the list are correctly identified.
func TestWhitelistHasTrue(t *testing.T) {
	w := NewWhitelist()
	w.Add("a")
	assert.True(t, w.Has("a"))
}

// assert that an empty whitelist means anything is allowed.
func TestWhitelistAllowedEmpty(t *testing.T) {
	w := NewWhitelist()
	assert.True(t, w.Allowed("a"))
}

// assert that a strict whitelist does not treat empty lists loosely.
func TestWhitelistAllowedStrict(t *testing.T) {
	w := NewWhitelist()
	w.Strict()
	assert.False(t, w.Allowed("a"))
}

// assert that only items in the whitelist are allowed.
func TestWhitelistAllowed(t *testing.T) {
	w := NewWhitelist()
	w.Add("a")
	assert.True(t, w.Allowed("a"))
	assert.False(t, w.Allowed("b"))
}

// assert that the size is accurately reported.
func TestWhitelistSize(t *testing.T) {
	w := NewWhitelist()
	assert.Zero(t, w.Size())
	w.Add("a")
	assert.Equal(t, 1, w.Size())
}

// assert that the strict method is chainable.
func TestWhitelistStrictChainable(t *testing.T) {
	w := NewWhitelist()
	assert.Equal(t, w, w.Strict())
}
