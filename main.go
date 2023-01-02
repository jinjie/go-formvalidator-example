package main

import (
	"fmt"
	"formvalidation/form"
)

func main() {
	// Setup the form with the data. For example, in Gin Gonic,
	// this can be something like ShouldBindJSON
	f := &form.UserForm{
		FirstName: "Jin Jie", // Pass Validation
		// FirstName: "Not Jin Jie", // Fail Validation
		LastName: "Kong   ", // Trimmed in SafeData
		Email:    "jinjie@example.com",
	}

	// Validates the form
	if err := form.Validate(f); err != nil {
		fmt.Println("Failed Validation")
		fmt.Println(err)
		return
	}

	// The form will be populated with the safe data if validates
	// successfully
	fmt.Println("Passed Validation")
	fmt.Printf("SafeData: %v\n", f)
}
