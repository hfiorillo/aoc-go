package main

import (
	"fmt"
	"io"
	"os"
)

func inputArray() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var perline int
	var nums []int
	for {
		_, err := fmt.Fscanf(file, "%d\n", &perline) // give a patter to scan
		if err != nil {
			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}
		nums = append(nums, perline)
	}
	// print out the nums array content

	// create chunks from input list
	chunk := 3
	var arr_chunks []int
	for i := 0; i < len(nums); i++ {
		chunks := nums[i:min(i+chunk, len(nums))]
		arr_chunks = append(arr_chunks, sum(chunks))
	}
	fmt.Println(arr_chunks)

	// count increases
	count := 0
	for i := 1; i < len(arr_chunks); i++ {
		if arr_chunks[i-1] < arr_chunks[i] {
			count += 1
		}
	}
	fmt.Println(count)
	return count
}

// main function
func main() {
	inputArray()
}

// check minimum
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// sum chunks together
func sum(chunks []int) int {
	sum := 0
	for _, v := range chunks {
		sum += v
	}
	return sum
}
