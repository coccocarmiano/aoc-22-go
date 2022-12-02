package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func minargmin(arr []int) (idx int, min int) {
	idx = 0
	min = 1e10

	for i, v := range arr {
		if v < min {
			min = v
			idx = i
		}
	}

	return idx, min
}

func maxargmax(arr []int) (idx int, max int) {
	idx = 0
	max = -1e10

	for i, v := range arr {
		if v > max {
			max = v
			idx = i
		}
	}

	return idx, max
}

func main() {
	file, err := os.Open("input.txt")
	const maxLineCapacity = 1024

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	top_3 := []int{0, 0, 0}
	tmp := 0

	for scanner.Scan() {
		line := scanner.Text()

		idx, min := minargmin(top_3)

		if tmp > min {
			top_3[idx] = tmp
		}

		if len(line) == 0 {
			tmp = 0
		} else {
			val, _ := strconv.Atoi(line) // Ignore error
			tmp += val
		}
	}

	_, max := maxargmax(top_3)
	fmt.Printf("Result I:%d\n", max)

	sum := 0
	for _, val := range top_3 {
		sum += val
	}
	fmt.Printf("Result II:%d\n", sum)

}
