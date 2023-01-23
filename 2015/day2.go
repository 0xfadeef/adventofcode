package main

import (
	"bufio"
	"fmt"
	"os"
)

func min(nums ...int) int {
	var m int
	for i, n := range(nums) {
		if n < m || i == 0 {
			m = n
		}
	}
	return m
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("missing input file\n")
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
	const pattern = "%dx%dx%d"

	var l, w, h int
	total := 0
	ribbon := 0

	for scanner.Scan() {
		line := scanner.Text()
		_, err := fmt.Sscanf(line, pattern, &l, &w, &h)

		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		s1, s2, s3 := l * w, w * h, h * l
		total += 2 * (s1 + s2 + s3) + min(s1, s2, s3)

		p1, p2, p3 := l + w, w + h, h + l
		ribbon += 2 * min(p1, p2, p3) + l * w * h
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Println(total, ribbon)
	}
}
