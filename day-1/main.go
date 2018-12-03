package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Starting...")

	dat, err := ioutil.ReadFile("input.txt")

	check(err)

	parts := strings.Split(string(dat), "\n")

	var result int64 = 0

	for i := 0; i < len(parts); i++ {
		value, err := strconv.ParseInt(parts[i][1:], 10, 32)

		check(err)

		if parts[i][0] == '+' {
			result += value
		} else {
			result -= value
		}
	}

	fmt.Println("Result: ", result)
}
