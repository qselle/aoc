package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func getItemPriority(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 96)
	}
	return int(c - 38)
}

func main() {
	input := flag.String("input", "input.txt", "aoc input to process")
	// part := flag.Int("part", 1, "aoc part (can be 1 or 2)")
	flag.Parse()

	f, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	sumDoubleItemPriority := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		currentDoubleItemPriority := 0
		start, end := line[:len(line)/2], line[len(line)/2:]
		for _, item := range start {
			if strings.ContainsRune(end, item) {
				currentDoubleItemPriority = getItemPriority(item)
				break
			}
		}
		sumDoubleItemPriority += currentDoubleItemPriority
	}
	fmt.Println(sumDoubleItemPriority)
}
