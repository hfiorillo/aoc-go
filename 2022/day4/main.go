package main

import (
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	"github.com/echojc/aocutil"

	"github.com/hfiorillo/aoc-go/helpers"
)

func main() {
	start := time.Now()
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}
	data, err := input.Strings(2022, 4)
	if err != nil {
		log.Fatal(err)
	}

	part1 := 0

	for _, items := range data {
		pair := strings.Split(items, ",")
		first, second := strings.Split(pair[0], "-"), strings.Split(pair[1], "-")
		switch {
		case helpers.ToIntList(first)[0] <= helpers.ToIntList(second)[0] && helpers.ToIntList(first)[1] >= helpers.ToIntList(second)[1]:
			part1++
		case helpers.ToIntList(first)[0] >= helpers.ToIntList(second)[0] && helpers.ToIntList(first)[1] <= helpers.ToIntList(second)[1]:
			part1++
		}
	}
	fmt.Println(part1)

	part2 := 0
	for _, items := range data {
		pair := strings.Split(items, ",")
		first, second := strings.Split(pair[0], "-"), strings.Split(pair[1], "-")
		max1 := math.Max(float64(helpers.ToIntList(first)[0]), float64(helpers.ToIntList(second)[0]))
		min2 := math.Min(float64(helpers.ToIntList(first)[1]), float64(helpers.ToIntList(second)[1]))
		if min2 >= max1 {
			fmt.Println(min2, max1, " | ", first, second)
			part2++
		}
	}
	fmt.Println(part2)
	fmt.Println(start)
}
