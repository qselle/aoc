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

func scoreCalcultationPart1(line string) int {
	game := strings.Split(line, " ")
	opponent, me := unify[game[0]], unify[game[1]]
	return battle[[2]int{opponent, me}] + me
}

var battleStrategy = map[[2]int]int{
	{rock, win}:      paper,
	{rock, draw}:     rock,
	{rock, lose}:     scissors,
	{paper, win}:     scissors,
	{paper, draw}:    paper,
	{paper, lose}:    rock,
	{scissors, win}:  rock,
	{scissors, draw}: scissors,
	{scissors, lose}: paper,
}

var strategyToDo = map[string]int{
	"X": lose,
	"Y": draw,
	"Z": win,
}

func scoreCalcultationPart2(line string) int {
	game := strings.Split(line, " ")
	opponent, needTo := unify[game[0]], strategyToDo[game[1]]
	return battleStrategy[[2]int{opponent, needTo}] + needTo
}

func main() {
	input := flag.String("input", "input.txt", "aoc input to process.")
	part := flag.Int("part", 1, "aoc exercice part (1 or 2).")
	flag.Parse()

	f, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	myScore := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		if *part == 1 {
			myScore += scoreCalcultationPart1(s.Text())
		} else {
			myScore += scoreCalcultationPart2(s.Text())
		}
	}
	if s.Err() != nil {
		panic(err)
	}
	fmt.Printf("My score is %d\n", myScore)
}
