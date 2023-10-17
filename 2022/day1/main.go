package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/echojc/aocutil"
	"github.com/hfiorillo/aoc-go/helpers"
)

// Convert list of strings into parts on "\n"
// Convert parts into list of ints
// Sum the parts
// Compare the parts

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

var elvish []int

func findMax(arr []int) (max int) {
	max = arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	data, err := input.Strings(2022, 1)
	if err != nil {
		log.Fatal(err)
	}

	list_of_elves := make([]int, 0)
	elvish = make([]int, 0)

	for i := 0; i < len(data); i++ {
		calories_int, err := helpers.ToInt(data[i])
		list_of_elves = append(list_of_elves, calories_int)
		if err != nil {
			elvish = append(elvish, sum(list_of_elves))
			list_of_elves = nil
		}
	}
	max_val := findMax(elvish)
	sort.Ints(elvish)
	max_3_total := elvish[len(elvish)-1] + elvish[len(elvish)-2] + elvish[len(elvish)-3]
	fmt.Println(elvish, max_val, max_3_total)
}
