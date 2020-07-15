package parser

import (
	"bufio"
	"log"
	"os"
	. "parzival/graphs-101/main/src/graph"
)

const PathError = "File not found according to given path: "

func BuildGraph(filename string) {
	graph := NewEmptyGraph()

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatalf(PathError + "%v", err)
		return
	}

	scanner := bufio.NewScanner(file)
	var totalVertices int

	for scanner.Scan() {
		line := scanner.Text()

		if TalksAboutVertices(line) {
			totalVertices, _ = Vertices(line)
			if TalksAboutVertices(line) {
				
			}
		}
	}
}
