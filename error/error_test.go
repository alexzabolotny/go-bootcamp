package main

import "testing"

func TestSqrtReturnError(t *testing.T) {
	_, err := Sqrt(-2)

	if err.Error() != "cannot Sqrt negative number: -2.000000" {
		t.Errorf("Error is not thrown, got: %s", err)
	}
}
