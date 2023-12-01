package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

func Part1(data []string) {
	total := 0
	for _, lines := range data {
		calibrationValues := []int{}
		for _, char := range lines {
			if num, err := strconv.Atoi(string(char)); err == nil {
				calibrationValues = append(calibrationValues, num)
			}
		}
		newValString := fmt.Sprintf("%d%d", calibrationValues[0], calibrationValues[len(calibrationValues)-1])
		newValInt, _ := strconv.Atoi(newValString)
		total += newValInt
	}
	fmt.Println("Part 1:", total)
}

func Part2(data []string) {

	doubles := map[string]string{
		"oneight":   "18",
		"twone":     "21",
		"threeight": "38",
		"fiveight":  "58",
		"sevenine":  "79",
		"eightwo":   "82",
		"eighthree": "83",
		"nineight":  "98",
	}

	singles := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	total := 0
	for _, lines := range data {

		for k, v := range doubles {
			lines = strings.Replace(lines, k, v, -1)
		}

		for k, v := range singles {
			lines = strings.Replace(lines, k, v, -1)
		}

		calibrationValues := []int{}
		for _, char := range lines {
			if num, err := strconv.Atoi(string(char)); err == nil {
				calibrationValues = append(calibrationValues, num)
			}
		}

		newValString := fmt.Sprintf("%d%d", calibrationValues[0], calibrationValues[len(calibrationValues)-1])
		newValInt, _ := strconv.Atoi(newValString)
		total += newValInt
	}

	fmt.Println("Part 2:", total)

}

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	data, err := input.Strings(2023, 1)
	if err != nil {
		log.Fatal(err)
	}
	Part1(data)
	Part2(data)
}
