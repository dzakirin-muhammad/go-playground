package main

import (
	"fmt"
	"os"
)

func main() {
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		// Running within GitHub Actions
		// Add your GitHub Actions specific code here
		fmt.Println("IN Github Action")
	} else {
		// Not running within GitHub Actions
		// Add your non-GitHub Actions code here
		fmt.Println("NOT Github Action")
	}
}
