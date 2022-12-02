package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	if err != nil {
		panic(err)
	}

	sm := [3][3]int{
		{4, 8, 3},
		{1, 5, 9},
		{7, 2, 6},
	}

	im := [3][3]int{
		{3, 4, 8},
		{1, 5, 9},
		{2, 6, 7},
	}

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	mapper := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"X": 0,
		"Y": 1,
		"Z": 2,
	}

	pts := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		him, you := split[0], split[1]
		nhim, nyou := mapper[him], mapper[you]
		pts += sm[nhim][nyou]
	}

	fmt.Printf("Result I: %d\n", pts)

	pts = 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		him, you := split[0], split[1]
		nhim, nyou := mapper[him], mapper[you]
		pts += im[nhim][nyou]
	}

	fmt.Printf("Result II: %d\n", pts)
}
