package tripcode_test

import (
	"fmt"

	"github.com/aquilax/tripcode"
)

func ExampleTripcode() {
	fmt.Println(tripcode.Tripcode("password"))
	// Output:
	// ozOtJW9BFA
}

func ExampleSecureTripcode() {
	fmt.Println(tripcode.SecureTripcode("password", "salt"))
	// Output:
	// Xtv3pggJ9U
}
