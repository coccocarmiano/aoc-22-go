package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	crates, moves := []string{}, []string{}

	sidx := 0
	for idx, line := range lines {
		if line == "" {
			sidx = idx
			break
		}
	}

	// Split file in two
	for idx := 0; idx < sidx; idx++ {
		crates = append(crates, lines[idx])
	}

	for idx := sidx + 1; idx < len(lines); idx++ {
		moves = append(moves, lines[idx])
	}

	triplets_re, _ := regexp.Compile("[0-9]+")

	triplets := [][3]int{}

	for _, line := range moves {
		matches := triplets_re.FindAllStringSubmatch(line, -1)
		_n, _from, _to := matches[0][0], matches[1][0], matches[2][0]
		from, _ := strconv.Atoi(_from)
		to, _ := strconv.Atoi(_to)
		n, _ := strconv.Atoi(_n)
		from -= 1
		to -= 1
		// DELETE THIS LINE
		n += 0

		move := [3]int{n, from, to}
		triplets = append(triplets, move)
	}

	// Get matrix width
	last_str := crates[len(crates)-1]
	re, _ := regexp.Compile("[0-9]+")
	reres := re.FindAllStringSubmatch(last_str, -1)
	w, _ := strconv.Atoi(reres[len(reres)-1][0])

	// Setup matrix
	char_re := regexp.MustCompile("[A-Z]")
	mat_one := make([][]string, w)
	mat_two := make([][]string, w)

	for idx := len(crates) - 1; idx >= 0; idx-- {
		line := crates[idx]
		res := char_re.FindAllStringSubmatchIndex(line, -1)
		for _, r := range res {
			c := line[r[0]:r[1]]
			x := r[0] / 4
			mat_one[x] = append(mat_one[x], c)
			mat_two[x] = append(mat_two[x], c)
		}
	}

	for _, move := range triplets {
		n, from, to := move[0], move[1], move[2]
		for i := 0; i < n; i++ {
			pop := mat_one[from][len(mat_one[from])-1]
			mat_one[from] = mat_one[from][:len(mat_one[from])-1]
			mat_one[to] = append(mat_one[to], pop)
		}
	}

	for _, move := range triplets {
		n, from, to := move[0], move[1], move[2]
		pop := mat_two[from][len(mat_two[from])-n : len(mat_two[from])]
		mat_two[from] = mat_two[from][:len(mat_two[from])-n]
		for _, p := range pop {
			mat_two[to] = append(mat_two[to], p)
		}
	}

	for _, line := range mat_one {
		print(line[len(line)-1])
	}
	println()

	for _, line := range mat_two {
		print(line[len(line)-1])
	}
	println()
}
