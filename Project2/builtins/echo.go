package builtins

import (
	"fmt"
	"strings"
)

func Echo(args ...string) {
	message := strings.Join(args, " ")
	fmt.Println(message)
}
