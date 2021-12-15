package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var horizontal int
	var depth int
	var aim int

	bytes, err := ioutil.ReadFile("day2_input.txt")

	if err != nil {
		return
	}

	data := string(bytes)
	lines := strings.Split(data, "\n")

outer:
	for _, line := range lines {
		var converted int
		var vals []string

		if vals = strings.Split(line, " "); len(vals) > 1 {
			converted, _ = strconv.Atoi(vals[1])

			switch vals[0] {
			case "forward":
				horizontal += converted
				depth += aim * converted
			case "down":
				aim += converted
			case "up":
				aim -= converted
			default:
				break outer
			}
		}
	}

	fmt.Println(horizontal * depth)
}
