package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	inputEnv        = "INPUT"
	defaultFilename = "input.txt"
)

func main() {
	filename := os.Getenv(inputEnv)
	if filename == "" {
		filename = defaultFilename
	}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if s.Err() != nil {
		panic(err)
	}

	var currentCalories, maxCalories uint64 = 0, 0
	var maxCaloriesElfID = 0
	for elfID, line := range lines {
		if line == "" {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
				maxCaloriesElfID = elfID
			}
			currentCalories = 0
			continue
		}
		foodCalories, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			panic(err)
		}
		currentCalories += foodCalories
	}
	fmt.Printf("The Elf carrying the most calories is the number %d with: %d calories.\n", maxCaloriesElfID, maxCalories)
}
