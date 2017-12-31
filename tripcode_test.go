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
			t.Errorf("Expected \"%s\", got \"%s\"", expected, salt)
		}
	}
}

func TestSecureTripcode(t *testing.T) {
	expect := "PqG4A0fkUs"
	trip := SecureTripcode("pass", "salt")
	if trip != expect {
		t.Errorf("SecureTripcode: expected \"%s\", got \"%s\"", expect, trip)
	}
}

func TestTripcode(t *testing.T) {
	cases := map[string]string{
		"asd":        "TAPy3blMsc",
		"adasd":      "IOuORdzMKw",
		"!@#$%^&*()": "BpZUCmJAIQ",
		"f}E":        "oUBoOTrysY",
		"©":          "", // Should be nothing?
		"訛":          "c8eDXvwFLQ",
		"!":          "KNs1o0VDv6",
		"@":          "z0MWdctOjE",
		"#":          "u2YjtUz8MU",
		"$":          "yflOPYrGcY",
		"%":          "1t98deumW.",
		"^":          "gBeeWo4hQg",
		"&":          "MhCJJ7GVT.",
		"*":          "o8gKYE6H8A",
		"(":          "SGn2Wwr9CY",
		")":          "E9k1wjKgHI",
		"-":          "tHbGiobWdM",
		"_":          "m3eoQIlU/U",
		"=":          "wmxP/NHJxA",
		"+":          "IHLbs/YhoA",
		"[":          "7h2f0/nQ3w",
		"]":          "rjM99frkZs",
		"{":          "odBt7a7lv6",
		"}":          "ATNP9hXHcg",
		";":          "zglc7ct1Ls",
		":":          ".BmRMKOub2",
		"'":          "8/08awL.AE",
		"\"":         "gt1azVccY2",
		"<":          "D1YGKrvmeg",
		">":          "afqVxck0Ts",
		",":          "YeQQgdCJE6",
		".":          "XONm83jaIU",
		"\\":         "9xUxYS2dlM",
		"?":          "cPUZU5OGFs",
		" ":          "wqLZLRuzPQ",
		"ññññ":       "",
		"糯ｫT弓(窶":     "Pants.f1Fk",
	}
	for pass, expected := range cases {
		trip := Tripcode(pass)
		if expected != trip {
			t.Errorf("Expected \"%s\", got \"%s\"", expected, trip)
		}
	}
}

func TestTripcodeOverflow(t *testing.T) {
	pass := ""
	for i := 0; i < 30000; i++ {
		pass += string(rune(i))
	}
	expected := "Ggih.96XPU"
	trip := Tripcode(pass)
	if expected != trip {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, trip)
	}
}
