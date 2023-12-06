package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	data, err := input.Strings(2023, 2)
	if err != nil {
		log.Fatal(err)
	}

	cubeMap := map[string]int{
		"blue":  14,
		"green": 13,
		"red":   12,
	}

	total := 0

	for i, games := range data {
		_, games, _ := strings.Cut(games, ": ")
		splitGames := strings.Split(games, "; ")
		gameCanBePlayed := true
		for _, set := range splitGames {
			pulls := strings.Split(set, ", ")
			for _, pull := range pulls {
				countStr, colour, _ := strings.Cut(pull, " ")
				countInt, _ := strconv.Atoi(countStr)
				if cubeMap[colour] < countInt {
					gameCanBePlayed = false
					break
				}
			}
		}
		if gameCanBePlayed {
			total += i + 1
		}
	}

	fmt.Println(total)
	total = 0

	// part 2

	for _, games := range data {
		cubeMap = map[string]int{
			"blue":  0,
			"green": 0,
			"red":   0,
		}
		_, games, _ := strings.Cut(games, ": ")
		splitGames := strings.Split(games, "; ")
		for _, set := range splitGames {
			pulls := strings.Split(set, ", ")
			for _, pull := range pulls {
				countStr, colour, _ := strings.Cut(pull, " ")
				countInt, _ := strconv.Atoi(countStr)
				if cubeMap[colour] < countInt {
					cubeMap[colour] = countInt
				}
			}
		}
		fmt.Println("cubeMap: ", cubeMap)
		total += cubeMap["blue"] * cubeMap["green"] * cubeMap["red"]
	}
	fmt.Println(total)
}
