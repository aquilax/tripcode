package tripcode

import (
	"testing"
)

func TestTripcode(t *testing.T) {
	cases := map[string]string{
		"asd":        "TAPy3blMsc",
		"adasd":      "IOuORdzMKw",
		"!@#$%^&*()": "BpZUCmJAIQ",
		"f}E":        "oUBoOTrysY",
		"©":          "",
		"訛":          "c8eDXvwFLQ",
		"'":          "8/08awL.AE",
	}
	for pass, expected := range cases {
		trip := Tripcode(pass)
		if expected != trip {
			t.Error("Expected "+expected+", got", trip)
		}
	}
}
