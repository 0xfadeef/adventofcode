package main

import (
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func is_valid(input []byte) bool {
	var c, prev byte
	var repeats = make(map[byte]bool)

	cur_seq_len := 1
	max_seq_len := cur_seq_len

	for n := 0; n < len(input); n++ {
		prev = c
		c = input[n]

		if c == 'i' || c == 'o' || c == 'l' {
			return false
		}
		if c == prev + 1 {
			cur_seq_len++
			continue
		} else {
			max_seq_len = max(max_seq_len, cur_seq_len)
			cur_seq_len = 1
		}
		if c == prev {
			repeats[c] = true
		}
	}
	return max_seq_len >= 3 && len(repeats) >= 2
}

func increment(input []byte) error {
	n := len(input)

next_char:
	if n--; n < 0 {
		return fmt.Errorf("Overflow!")
	}
	if c := input[n]; c < 'z' {
		input[n] = c + 1
		return nil
	} else {
		input[n] = 'a'
		goto next_char
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("missing input!\n")
		return
	}
	password := []byte(os.Args[1])

	for ; !is_valid(password); increment(password) {}
	fmt.Printf("%s\n", password)
}

