package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	OFFSET = 3
)

func determineCount(s []string) int64 {
	var count int64
	for _, val := range s {
		v, _ := strconv.ParseInt(val, 10, 0)
		count += v
	}
  return count
}

func main() {
	file, err := os.Open("day1_input.txt")

	if err != nil {
		fmt.Printf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	text := make([]string, 0)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

  var count int
	for i := 0; i+OFFSET < len(text); i++ {
    lhsCount := determineCount(text[i : i+OFFSET])
    rhsCount := determineCount(text[i+1 : i+1+OFFSET])

		if rhsCount > lhsCount {
			count++
		}
	}

	fmt.Println(count)
}
