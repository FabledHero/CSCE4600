package builtins

import (
	"bufio"
	"fmt"
	"os"
)

// Less reads a file and displays its contents in a paginated manner.
// Had AI help with this one. Gave most of the code, as i am new to Go and didn't know how to scan a file.
func Less(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		fmt.Printf("%4d  %s\n", lineNumber, scanner.Text())
		lineNumber++

		if lineNumber%20 == 0 {
			fmt.Print("Press Enter to continue...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return nil
}
