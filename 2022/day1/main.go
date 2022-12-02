package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Printf("Err: %s", err)
	}

	split := strings.Split(string(bytes), "\n")

	// Part 1

	//max := 0
	//current := 0
	//for _, s := range split {
	//if s == "" {
	//if current >= max {
	//max = current
	//}
	//current = 0
	//continue
	//}
	//i, err := strconv.ParseInt(s, 10, strconv.IntSize)
	//if err != nil {
	//fmt.Println(err)
	//}
	//current += int(i)
	//}

	// Part 2
	totals := []int{}
	current := 0
	for _, s := range split {
		if s == "" {
			totals = append(totals, current)
			current = 0
			continue
		}
		i, err := strconv.ParseInt(s, 10, strconv.IntSize)
		if err != nil {
			fmt.Println(err)
		}
		current += int(i)
	}
	sort.Slice(totals, func(a, b int) bool {
		return totals[a] > totals[b]
	})

	sum := 0
	for _, n := range totals[:3] {
		sum += n
	}

	fmt.Println(sum)
}
