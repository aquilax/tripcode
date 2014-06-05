package tripcode

import (
	"testing"
)

func TestGenerateSalt(t *testing.T) {
	cases := map[string]string{
		"asd":        "sd",
		"adasd":      "da",
		"!@#$%^&*()": "G.",
		"f}E":        ".E",
		"©":          "H.",
		"訛":          "H.",
		"'":          "H.",
	}
	for pass, expected := range cases {
		salt := generateSalt(pass)
		if expected != salt {
			t.Error("Expected "+expected+", got", salt)
		}
	}

}

func TestTripcode(t *testing.T) {
	cases := map[string]string{
		"asd":        "TAPy3blMsc",
		"adasd":      "IOuORdzMKw",
		"!@#$%^&*()": "BpZUCmJAIQ",
		"f}E":        "oUBoOTrysY",
		"©":          "",
		"訛":          "c8eDXvwFLQ",
		// @TODO: Figure out this case
		//"'":          "8/08awL.AE",
	}
	for pass, expected := range cases {
		trip := Tripcode(pass)
		if expected != trip {
			t.Error("Expected "+expected+", got", trip)
		}
	}
}
