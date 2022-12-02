package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	Number   uint64
	Calories uint64
}

func main() {
	input := flag.String("input", "input.txt", "input to process")
	limit := flag.Int("limit", 1, "elf limit to print (by descending order)")
	flag.Parse()

	f, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	var currentCalories, elfNumber uint64 = 0, 0
	var elfes []Elf

	for s.Scan() {
		line := s.Text()
		if line == "" {
			elfes = append(elfes, Elf{elfNumber + 1, currentCalories})
			currentCalories = 0
			elfNumber += 1
			continue
		}
		foodCalories, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			panic(err)
		}
		currentCalories += foodCalories
	}
	if s.Err() != nil {
		panic(err)
	}

	sort.SliceStable(elfes, func(i, j int) bool {
		return elfes[i].Calories > elfes[j].Calories
	})
	var total uint64
	for num, elf := range elfes {
		if num == *limit {
			break
		}
		fmt.Printf("I'm %d and I'm carrying %d\n", elf.Number, elf.Calories)
		total += elf.Calories
	}
	fmt.Printf("%d calories carrying by the top %d elf(es)\n", total, *limit)
}
