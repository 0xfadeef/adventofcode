package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	TURNON = iota
	TURNOFF
	TOGGLE
)

var actions = []string{TURNON: "turn on", TURNOFF: "turn off", TOGGLE: "toggle"}
var lights [1000][1000]uint

func cutActionPrefix(s string, action *uint) (after string, ok bool) {
	after = s
	for n, v := range actions {
		if after, ok = strings.CutPrefix(s, v); ok {
			*action = uint(n)
			return
		}
	}
	return
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

	var count = 0
	var action, x1, y1, x2, y2 uint

	for scanner.Scan() {
		line := scanner.Text()
		tail, ok := cutActionPrefix(line, &action)

		if ok {
			fmt.Sscanf(tail, "%d,%d through %d,%d", &x1, &y1, &x2, &y2)
		} else {
			fmt.Printf("invalid input: %s\n", line)
			return
		}

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				before := lights[x][y]

				switch action {
				case TURNON:
					lights[x][y] = 1
				case TURNOFF:
					lights[x][y] = 0
				case TOGGLE:
					lights[x][y] = before ^ 1
				}
				count += int(lights[x][y] - before)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Println(count)
	}
}
