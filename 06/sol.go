package main

import (
	"bufio"
	"os"
)

func all_different(w string) bool {
	for i, c := range w {
		for j, d := range w {
			if i != j && c == d {
				return false
			}
		}
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	wone := []string{}
	wtwo := []string{}

	for i := 0; i < len(line)-3; i++ {
		tmp := line[i : i+4]
		wone = append(wone, tmp)
	}

	for i := 0; i < len(line)-13; i++ {
		tmp := line[i : i+14]
		wtwo = append(wtwo, tmp)
	}

	for idx, w := range wone {
		if all_different(w) {
			println("Result I: ", idx+4)
			break
		}
	}

	for idx, w := range wtwo {
		if all_different(w) {
			println("Result II: ", idx+14)
			break
		}
	}

}
