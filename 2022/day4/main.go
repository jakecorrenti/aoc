package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	A int
	B int
}

func initPair(s string) Pair {
	ends := strings.Split(s, "-")
	a, _ := strconv.Atoi(string(ends[0]))
	b, _ := strconv.Atoi(string(ends[1]))

	return Pair{
		A: a,
		B: b,
	}
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}

	split := strings.Split(string(bytes), "\n")
	total := 0
	for _, s := range split {
		if len(s) <= 0 {
			continue
		}
		pair := strings.Split(s, ",")
		pairA := initPair(pair[0])
		pairB := initPair(pair[1])

		// Part 1
		//if (pairB.A >= pairA.A && pairB.B <= pairA.B) || (pairA.A >= pairB.A && pairA.B <= pairB.B) {
		//total += 1
		//}

		// Part 2
		// checks fully contained by one or the other
		if (pairB.A >= pairA.A && pairB.B <= pairA.B) || (pairA.A >= pairB.A && pairA.B <= pairB.B) {
			total += 1
			continue
		}

		// checks partial contianment
		if (pairA.A <= pairB.A && pairB.A <= pairA.B) || (pairB.A <= pairA.A && pairA.A <= pairB.B) {
			total += 1
		}

	}
	fmt.Println(total)
}
