package hello

import (
	"strings"
)

func SayHello(names []string) string {
	if len(names) == 0 {
		names = []string{"World"}
	}
	return "Hello, " + strings.Join(names, ", ") + "!"
}
