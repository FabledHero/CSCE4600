package builtins

import (
	"fmt"
	"os"
	"errors"
)

func Pwd() error {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		//error catching
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(currentDir)
	return nil
}
