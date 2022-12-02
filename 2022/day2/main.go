package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bytes)
}
