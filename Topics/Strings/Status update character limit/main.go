package main

import (
	"bufio"
	"fmt"
	"os"
)

//nolint:gomnd // <-- DO NOT delete this comment!
func main() {
	// DO NOT delete or modify the code block below!
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	status := scanner.Text()

	// Count the characters in the `status` below:
	characterCount := len(status)

	// Check if `characterCount` is within the limit of 140 characters
	// And print the appropriate message:
	if characterCount <= 140 {
		fmt.Println("Status is within the 140 character limit")
	} else {
		fmt.Println("Status exceeds the 140 character limit")
	}
}
