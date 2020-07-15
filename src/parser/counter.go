package parser

import (
	"strconv"
	"strings"
)

// Checks if the parsing is at the first line.
func TalksAboutVertices(givenLine string) bool {
	return strings.Contains(givenLine, "*vertices")
}

// Checks if the parsing reached the middle of
// the .net file.
func TalksAboutEdges(givenLine string) bool {
	return strings.Contains(givenLine, "*edges")
}

// Gets the total vertices in the graph according
// to the .net passed as argument.
func Vertices(firstLine string) (int, error) {
	toBeDivided := strings.SplitAfter(firstLine, " ")
	divided := toBeDivided[0]
	return strconv.Atoi(divided)
}
