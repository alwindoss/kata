package utils

import "bufio"

// FindUniqueLines finds the unique lines from the given input
func FindUniqueLines(input *bufio.Scanner) (map[string]int, []string) {
	counts := make(map[string]int)
	var orderedUniqueKeyList []string
	for input.Scan() {
		t := input.Text()
		counts[t]++
		if counts[t] == 1 {
			orderedUniqueKeyList = append(orderedUniqueKeyList, t)
		}
	}
	return counts, orderedUniqueKeyList
}
