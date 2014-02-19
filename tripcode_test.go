package tripcode

import (
"testing"
)
func TestTripcode(t *testing.T) {
	pass := "asd"
	expected := "TAPy3blMsc"
	trip := Tripcode(pass)
	if expected != trip {
		t.Error("Expected , got ", trip)
	}
}
