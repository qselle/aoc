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
	part := flag.Int("part", 1, "aoc part (can be 1 or 2)")
	flag.Parse()

	f, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	if *part == 1 {
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
	} else if *part == 2 {
		index := 0
		sum := 0
		s := bufio.NewScanner(f)
		var group []string
		for s.Scan() {
			index++
			line := s.Text()
			group = append(group, line)
			if index%3 == 0 {
				for _, item := range group[0] {
					if strings.ContainsRune(group[1], item) {
						if strings.ContainsRune(group[2], item) {
							sum += getItemPriority(item)
							break
						}
					}
				}
				group = nil
			}
		}
		fmt.Println(sum)
	} else {
		fmt.Println("wrong part dumbass")
	}
}
