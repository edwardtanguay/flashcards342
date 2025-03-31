package main

import (
	"fmt"

	"github.com/edwardtanguay/flashcards342/utils"
)

func main() {
	lines := utils.GetLinesFromFile("test.txt")
	fmt.Println(lines)
}