package main

import (
	"fmt"
	"log"
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
	for _, items := range data {
		pairs := strings.Split(items, ",")
		first, second := strings.Split(pairs[0], "-"), strings.Split(pairs[1], "-")
		fmt.Println(first, helpers.ToIntList(first)[0], helpers.ToIntList(first)[1])
		fmt.Println(second, helpers.ToIntList(second)[0], helpers.ToIntList(second)[1])
		for i := helpers.ToIntList(first)[0]; i <= helpers.ToIntList(first)[1]; i++ {
			for j := helpers.ToIntList(second)[0]; j <= helpers.ToIntList(second)[1]; j++ {
				fmt.Println(i, j)
			}
		}
	}
	fmt.Println(time.Since(start))
}

// Input: [... 5-79,6-6 95-95,19-95 37-94,38-94 2-12,2-36 9-46,46-47 58-80,66-81]
// Item: strings 5-79,6-6
// Sort into pairs, convert to integers
// Create array of numbers
// Is the first number in 1 before / after the first number in 2
// Is the second number in 1 before / after the second number in 2
// How many times do the 2 arrays ints overlap fully (1-2) (1-3)
