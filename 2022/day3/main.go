package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/echojc/aocutil"
)

func findPriority(n rune) int {
	var priotity int
	if string(n) >= "A" && string(n) <= "Z" {
		priotity = int(n) - 38
		return priotity
	}
	if string(n) >= "a" && string(n) <= "z" {
		priotity = int(n) - 96
		return priotity
	}
	return 0
}

func main() {
	start := time.Now()
	total_1 := 0
	total_2 := 0
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	data, err := input.Strings(2022, 3)
	if err != nil {
		log.Fatal(err)
	}

	groups := make([]string, 0)

	for _, items := range data {
		first := items[:len(items)/2]
		second := items[len(items)/2:]

		for _, g := range first {
			if strings.ContainsRune(second, g) {
				total_1 += findPriority(g)
				break
			}
		}

		groups = append(groups, items)
		if len(groups) == 3 {
			for _, r := range groups[0] {
				if strings.ContainsRune(groups[1], r) && strings.ContainsRune(groups[2], r) {
					total_2 += findPriority(r)
					break
				}
			}
			groups = make([]string, 0)
		}
	}
	fmt.Printf("Part 1: %d Part 2: %d\n", total_1, +total_2)
	fmt.Println(time.Since(start))
}
