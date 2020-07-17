package test

import "testing"

func AssertEqualInt(t *testing.T, want, got int) {
	t.Helper()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func AssertEqualString(t *testing.T, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
