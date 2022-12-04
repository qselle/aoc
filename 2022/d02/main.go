package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	rock = iota + 1
	paper
	scissors
	lose = 0
	draw = 3
	win  = 6
)

var unify = map[string]int{
	"A": rock,
	"B": paper,
	"C": scissors,
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

var battle = map[[2]int]int{
	{rock, rock}:         draw,
	{rock, paper}:        win,
	{rock, scissors}:     lose,
	{paper, rock}:        lose,
	{paper, paper}:       draw,
	{paper, scissors}:    win,
	{scissors, rock}:     win,
	{scissors, paper}:    lose,
	{scissors, scissors}: draw,
}

func main() {
	input := flag.String("input", "input.txt", "aoc input to process.")
	flag.Parse()

	f, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	myScore := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		game := strings.Split(line, " ")
		opponent, me := unify[game[0]], unify[game[1]]
		myScore += battle[[2]int{opponent, me}] + me
	}
	if s.Err() != nil {
		panic(err)
	}
	fmt.Printf("My score is %d\n", myScore)
}
