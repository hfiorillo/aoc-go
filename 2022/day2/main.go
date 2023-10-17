package main

import (
	"fmt"
	"log"
	"time"

	"github.com/echojc/aocutil"
)

func main() {
	start := time.Now()
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	data, err := input.Strings(2022, 2)
	if err != nil {
		log.Fatal(err)
	}

	scores := map[string]struct{ p1, p2 int }{
		"A X": {1 + 3, 3 + 0}, "A Y": {2 + 6, 1 + 3}, "A Z": {3 + 0, 2 + 6},
		"B X": {1 + 0, 1 + 0}, "B Y": {2 + 3, 2 + 3}, "B Z": {3 + 6, 3 + 6},
		"C X": {1 + 6, 2 + 0}, "C Y": {2 + 0, 3 + 3}, "C Z": {3 + 3, 1 + 6},
	}

	part1, part2 := 0, 0

	for _, games := range data {
		part1 += scores[games].p1
		part2 += scores[games].p2
	}
	fmt.Println(part1, part2)
	fmt.Println(time.Since(start))
}
