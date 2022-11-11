package assert

import "testing"

// Equal checks if two values are equal else fails the test
func Equal[T comparable](t *testing.T, expected T, v T) {
	if expected != v {
		t.Errorf("expected %#v but got %#v", expected, v)
	}
}

// False checks whether v is false
func False(t *testing.T, v bool) {
	Equal(t, false, v)
}

// True checks whether v is true
func True(t *testing.T, v bool) {
	Equal(t, true, v)
}
