package utils

import (
	"fmt"
	"io"
)

// PrintStringCount prints the string count to a io.Writer os.Stdout by default
func PrintStringCount(orderedUniqueKeyList []string, input map[string]int, out io.Writer) {
	for _, line := range orderedUniqueKeyList {
		fmt.Fprintf(out, "%d\t%s\n", input[line], line)
	}
}
