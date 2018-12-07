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

func parseClaim(claim string) (int, int, int, int) {
	parts := strings.Split(claim, " ")

	dimensions := strings.Split(parts[3], "x")

	width, err := strconv.Atoi(dimensions[0])

	check(err)

	height, err := strconv.Atoi(dimensions[1])

	check(err)

	position := strings.Split(parts[2][:len(parts[2])-1], ",")

	left, err := strconv.Atoi(position[0])

	check(err)

	top, err := strconv.Atoi(position[1])

	check(err)

	return width, height, top, left
}

func main() {
	const FABRIC_WIDTH = 1000
	const FABRIC_HEIGHT = 1000

	dat, err := ioutil.ReadFile("input.txt")

	check(err)

	claims := strings.Split(string(dat), "\n")

	var fabric [FABRIC_WIDTH][FABRIC_HEIGHT]int

	for _, claim := range claims {
		width, height, top, left := parseClaim(claim)

		for i := 0; i < width; i++ {
			for j := 0; j < height; j++ {
				fabric[left+i][top+j] += 1
			}
		}
	}

	overlaps := 0

	for i := 0; i < FABRIC_WIDTH; i++ {
		for j := 0; j < FABRIC_HEIGHT; j++ {
			if fabric[i][j] > 1 {
				overlaps += 1
			}
		}
	}

	fmt.Println("Number of overlaps: ", overlaps)
}
