package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// A = rock, B = paper, C = scissors
// X = rock, Y = paper, Z = scissors
// X = loss, Y = draw, Z = win

// sum of your score for round + outcome of the round
// 1 = rock, 2 = paper, 3 = scissors
// 0 = loss, 3 = draw, 6 = win

// rock > scissors
// paper > rock
// scissors > paper
const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3

	LOSS = 0
	DRAW = 3
	WIN  = 6
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	split := strings.Split(string(bytes), "\n")
	runningScore := 0

	for _, s := range split[:len(split)-1] {
		s1 := s[0]
		s2 := s[2]
		switch string(s1) {
		case "A":
			switch string(s2) {
			case "X":
				runningScore += LOSS + SCISSORS
			case "Y":
				runningScore += DRAW + ROCK
			default:
				runningScore += WIN + PAPER
			}
		case "B":
			switch string(s2) {
			case "X":
				runningScore += LOSS + ROCK
			case "Y":
				runningScore += DRAW + PAPER
			default:
				runningScore += WIN + SCISSORS
			}
		case "C":
			switch string(s2) {
			case "X":
				runningScore += LOSS + PAPER
			case "Y":
				runningScore += DRAW + SCISSORS
			default:
				runningScore += WIN + ROCK
			}
		}
	}
	fmt.Println(runningScore)
}
