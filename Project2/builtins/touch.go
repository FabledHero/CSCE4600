package builtins

import (
	"fmt"
	"os"
)

// Touch creates an empty file or updates the access and modification times of a file.
func Touch(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File created/updated:", filename)
	}
	file.Close()
	return nil
}
