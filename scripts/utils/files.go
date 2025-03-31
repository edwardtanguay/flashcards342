package utils

import (
	"os"
	"strings"
)

func GetLinesFromFile(fileName string) []string {
	byteContents, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	contents := string(byteContents)
	lines := strings.Split(contents, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)	
	}
	return lines
}