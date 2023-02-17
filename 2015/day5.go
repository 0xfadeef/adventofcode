package main

import (
	"bufio"
	"fmt"
	"os"
)

var vowels = []byte{'a', 'e', 'i', 'o', 'u'}
var forbidden = []string{"ab", "cd", "pq", "xy"}

func contains[T comparable](elems []T, v T) bool {
	for _, e := range elems {
		if e == v {
			return true
		}
	}
	return false
}

func is_nice(line string) bool {
	vowel_count := 0
	double_letter := false
	last_index := len(line) - 1

	for i := 0; i < last_index; i++ {
		if contains(forbidden, line[i:i+2]) {
			return false
		}
		if contains(vowels, line[i]) {
			vowel_count++
		}
		if line[i] == line[i+1] {
			double_letter = true
		}
	}
	if contains(vowels, line[last_index]) {
		vowel_count++
	}
	return vowel_count >= 3 && double_letter
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing input file")
		return
	}

	input := os.Args[1]
	f, err := os.Open(input)

	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var nice = 0

	for scanner.Scan() {
		line := scanner.Text()
		if is_nice(line) {
			nice++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Println(nice)
	}
}
