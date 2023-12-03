package builtins

import (
	"fmt"
	"strings"
	"errors"
)

func Echo(args ...string) error {
	message := strings.Join(args, " ")
	fmt.Println(message)
	return nil
}
