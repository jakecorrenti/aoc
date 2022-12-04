package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	split := strings.Split(string(bytes), "\n")

	shared := []rune{}

	// Part 1

	//for _, s := range split {
	//if len(s) == 0 {
	//break
	//}
	//s1 := s[:len(s)/2]
	//s2 := s[len(s)/2:]
	//for _, c := range s1 {
	//if strings.Contains(s2, string(c)) {
	//shared = append(shared, c)
	//break
	//}
	//}
	//}

	//total := 0
	//for _, s := range shared {
	//if unicode.IsLower(s) {
	//total += int(s-'a') + 1
	//} else {
	//total += int(s-'A') + 27
	//}
	//}
	//fmt.Println(total)

	// Part 2
	for i := 0; i < len(split)-3; i += 3 {
		s1 := split[i]
		s2 := split[i+1]
		s3 := split[i+2]
		for _, c := range s1 {
			if strings.Contains(s2, string(c)) && strings.Contains(s3, string(c)) {
				shared = append(shared, c)
				break
			}
		}
	}

	total := 0
	for _, s := range shared {
		if unicode.IsLower(s) {
			total += int(s-'a') + 1
		} else {
			total += int(s-'A') + 27
		}
	}
	fmt.Println(total)
}
