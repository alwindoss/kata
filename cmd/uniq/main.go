package main

import (
	"bufio"
	"os"

	"github.com/alwindoss/kata/internal/utils"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	counts, orderedUniqueKeyList := utils.FindUniqueLines(input)
	utils.PrintStringCount(orderedUniqueKeyList, counts, os.Stdout)
}
