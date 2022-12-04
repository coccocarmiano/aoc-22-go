package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic("File not found")
	}

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	cnt_one := 0
	cnt_two := 0

	for _, line := range lines {
		split := strings.Split(line, ",")
		left, right := split[0], split[1]

		lsplit := strings.Split(left, "-")
		ll, lr := lsplit[0], lsplit[1]

		rsplit := strings.Split(right, "-")
		rl, rr := rsplit[0], rsplit[1]

		lstart, _ := strconv.Atoi(ll)
		lend, _ := strconv.Atoi(lr)
		rstart, _ := strconv.Atoi(rl)
		rend, _ := strconv.Atoi(rr)

		lmap, rmap := make(map[int]bool), make(map[int]bool)

		for i := lstart; i <= lend; i++ {
			lmap[i] = true
		}

		for i := rstart; i <= rend; i++ {
			rmap[i] = true
		}

		cond_one := true
		cond_two := true
		cond_three := false

		for k := range lmap {
			if !rmap[k] {
				cond_one = false
				break
			}
		}

		for k := range rmap {
			if !lmap[k] {
				cond_two = false
				break
			}
		}

		for k := range lmap {
			if rmap[k] {
				cond_three = true
				break
			}
		}

		if cond_one || cond_two {
			cnt_one++
		}

		if cond_three {
			cnt_two++
		}
	}

	fmt.Printf("Result I: %d\n", cnt_one)
	fmt.Printf("Result II: %d\n", cnt_two)

}
