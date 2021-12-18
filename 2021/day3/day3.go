package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {

	bytes, err := ioutil.ReadFile("day3_input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	data := string(bytes)
	splits := strings.Split(data, "\n")
	splits = splits[:len(splits)-1]

	var gamma float64
	var epsilon float64

	for i := 0; i < len(splits[0]); i++ {
		frequency := make(map[int]int, 2)

		for _, s := range splits {
			conv, err := strconv.Atoi(string(s[i]))

			if err != nil {
				log.Fatal(err)
				return
			}

			frequency[conv] += 1
		}

		if frequency[1] > frequency[0] {
			gamma += math.Pow(2, float64(len(splits[0])-i)-1)
		} else {
			epsilon += math.Pow(2, float64(len(splits[0])-i)-1)
		}
	}

  fmt.Printf("gamma: %f\n", gamma)
  fmt.Printf("epsilon: %f\n", epsilon)
  fmt.Printf("solution: %f\n", gamma * epsilon)

  // life support rating = oxygen generator rating * co2 scrubber rating
}
