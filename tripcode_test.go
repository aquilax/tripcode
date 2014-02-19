package tripcode

import (
	"testing"
)

func TestTripcode(t *testing.T) {
	cases := map[string]string{
		"asd":        "TAPy3blMsc",
		"adasd":      "IOuORdzMKw",
		"!@#$%^&*()": "96TA4mR8Fc",
	}
	for pass, expected := range cases {
		trip := Tripcode(pass)
		if expected != trip {
			t.Error("Expected "+expected+", got", trip)
		}
	}
}
