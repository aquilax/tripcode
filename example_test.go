package tripcode_test

import (
	"fmt"

	"github.com/aquilax/tripcode"
)

func ExampleTripcode_Tripcode() {
	fmt.Println(tripcode.Tripcode("password"))
	// Output:
	// ozOtJW9BFA
}

func ExampleTripcode_SecureTripcode() {
	fmt.Println(tripcode.SecureTripcode("password", "salt"))
	// Output:
	// Xtv3pggJ9U
}
