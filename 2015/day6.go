package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const toggle_switch = false

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

func execAction(action uint, light *uint) int {
	before := *light
	switch action {
	case TURNON:
		if toggle_switch {
			*light = 1
		} else {
			*light = before + 1
		}
	case TURNOFF:
		if toggle_switch {
			*light = 0
		} else if before > 0 {
			*light = before - 1
		}
	case TOGGLE:
		if toggle_switch {
			*light = before ^ 1
		} else {
			*light = before + 2
		}
	}
	return int(*light - before)
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
				count += execAction(action, &lights[x][y])
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Println(count)
	}
}
