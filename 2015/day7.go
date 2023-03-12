package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


type Signal uint16
var table = make(map[string]string)
var cache = make(map[string]Signal)

func signal(wire string) Signal {
	if cached, ok := cache[wire]; ok {
		return cached
	}
	source, ok := table[wire]
	if !ok {
		panic(fmt.Errorf("missing signal source for wire \"%s\"", wire))
	}

	var sig Signal
	parts := strings.Split(source, " ")

	switch len(parts) {
	case 1:
		sig = eval(parts[0])
	case 2:
		if parts[0] != "NOT" {
			panic(fmt.Errorf("unknown logic gate: %s", source))
		}
		sig = ^eval(parts[1])
	case 3:
		switch parts[1] {
		case "AND":
			sig = eval(parts[0]) & eval(parts[2])
		case "OR":
			sig = eval(parts[0]) | eval(parts[2])
		case "LSHIFT":
			sig = eval(parts[0]) << eval(parts[2])
		case "RSHIFT":
			sig = eval(parts[0]) >> eval(parts[2])
		default:
			panic(fmt.Errorf("unknown logic gate: %s", source))
		}
	default:
		panic(fmt.Errorf("invalid singal source for wire \"%s\"", wire))
	}
	cache[wire] = sig
	return sig
}

func eval(value string) Signal {
	var sig Signal
	n, _ := fmt.Sscanf(value, "%d", &sig)
	if n == 1 {
		return sig
	} else {
		return signal(value)
	}
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

	var source, wire string

	for scanner.Scan() {
		line := scanner.Text()
		subs := strings.Split(line, " -> ")
		source, wire = subs[0], subs[1]

		if _, ok := table[wire]; ok {
			fmt.Printf("wire %s gets singal from more than one source!\n", wire)
			return
		}
		// we wont parse the source string until it's actually required
		table[wire] = source
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	a_sig := signal("a")
	fmt.Println(a_sig)
}

