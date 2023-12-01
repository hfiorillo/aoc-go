package main

import (
	_ "embed"
	"fmt"
	"strings"
)

// Data looks like...
// [D]
// [N] [C]
// [Z] [M] [P]
// 1   2   3

// Return: [['Z', 'N'], ['M', 'C', 'D'], ['P']]

//go:embed 2022_5.txt
var input string

func init() {
	tippingPoint := strings.Index(input, "/n")
	columns := input[:tippingPoint+1]
	// _ := input[tippingPoint:]
	fmt.Println(columns)
}

func main() {
	// start := time.Now()
	// input, err := aocutil.NewInputFromFile("session_id")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// data, err := input.Strings(2022, 5)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// capture the starting positions for each crane (9)
	// sort the crates in the correct crane

	// cranes := make([][]string, 9)

	// init()
}

// This is key and needs to be extensible
func getCrateChar(crate string) string {
	for _, v := range crate {
		// if the character is not [ ] or then append it to a list and
		// if strings.ContainsRune(unsorted, v) {
		// }
		fmt.Println(v)
	}
	return ""
}

func getData(data string) (string, string) {
	tippingPoint := strings.Index(data, "/nmove")
	columns := data[:tippingPoint]
	instructions := data[tippingPoint:]
	return columns, instructions
}

func readCrate(crate []string) []string {
	return []string{""}
}

func readMovement(string) (int, int) {
	return 0, 0
}

func movementFunc(crate1, crate2 []string, amount int) ([]string, []string) {
	return []string{}, []string{}
}
