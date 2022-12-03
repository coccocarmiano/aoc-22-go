package main

import (
	"bufio"
	"fmt"
	"os"
)

// lettere = tipo di oggetti
// 1 riga = 1 zaino
// metà riga = compartimento

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	result := 0
	result_two := 0

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	map_three := make(map[byte]int)

	for idx, line := range lines {
		div := len(line) / 2
		str_one, str_two := line[:div], line[div:]
		map_one, map_two := make(map[rune]bool), make(map[rune]bool)

		for _, c := range str_one {
			map_one[c] = true
		}

		for _, c := range str_two {
			map_two[c] = true
		}

		for k := range map_one {
			if _, ok := map_two[k]; ok {
				result += ctoi(byte(k))
			}
		}

		tmp := map[byte]bool{}
		for _, c := range line {
			if _, ok := tmp[byte(c)]; !ok {
				map_three[byte(c)]++
				tmp[byte(c)] = true
			}
		}

		if (idx+1)%3 == 0 {
			for k, v := range map_three {
				if v == 3 {
					result_two += ctoi(byte(k))
				}
			}
			map_three = make(map[byte]int)
		}
	}

	fmt.Printf("Result I: %d\n", result)
	fmt.Printf("Result II: %d\n", result_two)
}

func ctoi(c byte) int {
	c = c - 64
	switch {
	case c < 27:
		// A..B -> a..b
		return int(c + 26)
	case c > 32:
		return int(c - (26 + 6))
	}
	panic("invalid character")
}
