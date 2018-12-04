package main

import (
	"fmt"
	"io/ioutil"
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

	codes := strings.Split(string(dat), "\n")

	twos := 0
	threes := 0

	for i := 0; i < len(codes); i++ {
		occurs := make(map[rune]int)

		code := codes[i]

		for _, alpha := range code {
			occurs[alpha] += 1
		}

		count_two := 0
		count_three := 0

		for key := range occurs {
			if occurs[key] == 2 {
				count_two = 1
			} else if occurs[key] == 3 {
				count_three = 1
			}
		}

		twos += count_two
		threes += count_three
	}

	fmt.Println("Checksum: ", twos*threes)
}
