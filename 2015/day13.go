package main

import (
	"bufio"
	"fmt"
	"os"
)


type Pair struct {
	guest1, guest2 int
}

var happiness = make(map[Pair]int)
var guests = make([]string, 0, 32)

func get_guest_index(guest string) (k int) {
	for k = 0; k < len(guests); k++ {
		if guest == guests[k] {
			return
		}
	}
	guests = append(guests, guest)
	return
}


/* This is implementation of Held-Karp algorithm for TSP. */

/* Compute the lexicographically next bit permutation
   https://graphics.stanford.edu/~seander/bithacks.html#NextBitPermutation
*/
func next(v uint) uint {
	t := (v | (v - 1)) + 1
	w := t | ((((t & -t) / (v & -v)) >> 1) - 1)
	return w
}

func hamilton_cycle_max_length(n int, w map[Pair]int) int {
	n--  // exclude starting point

	/* To represent the set of nodes we use bitmask, where
	   bit position corresponds to the index of the node.
	*/
	var mask  uint
	var limit uint = 1 << n

	type Set struct {
		mask uint
		e int
	}
	var g = make(map[Set]int)

	compute_longest_path := func (set Set) int {
		mask := set.mask &^ (1 << set.e)
		opt := -1 << 31

		for m := 0; m < n; m++ {
			if mask & (1 << m) == 0 {
				continue
			}
			d := g[Set{mask, m}] + w[Pair{m, set.e}]
			opt = max(opt, d)
		}
		return opt
	}

	for k := 0; k < n; k++ {
		mask = 1 << k
		g[Set{mask, k}] = w[Pair{n, k}]
	}

	for s := 2; s <= n; s++ {
		for mask = (1 << s) - 1; mask < limit; mask = next(mask) {
			for k := 0; k < n; k++ {
				if mask & (1 << k) == 0 {
					continue
				}
				set := Set{mask, k}
				g[set] = compute_longest_path(set)
			}
		}
	}

	mask = (1 << n) - 1
	opt := -1 << 31

	for k := 0; k < n; k++ {
		d := g[Set{mask, k}] + w[Pair{k, n}]
		opt = max(opt, d)
	}
	return opt
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

	const pattern = "%s would %s %d happiness units by sitting next to %s."
	scanner := bufio.NewScanner(f)

	var guest1, guest2, verb string
	var happy_units int

	for scanner.Scan() {
		bytes := scanner.Bytes()
		n := len(bytes)

		line := string(bytes[:n-1])  // remove trailing period
		fmt.Sscanf(line, pattern, &guest1, &verb, &happy_units, &guest2)

		if verb == "lose" {
			happy_units = -happy_units
		} else if verb != "gain" {
			fmt.Printf("error: %s\n", line)
			return
		}

		g1 := get_guest_index(guest1)
		g2 := get_guest_index(guest2)

		pair1 := Pair{g1, g2}
		pair2 := Pair{g2, g1}

		if h, ok := happiness[pair2]; ok {
			happy_units += h
			happiness[pair2] = happy_units
		}
		happiness[pair1] = happy_units
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	total := hamilton_cycle_max_length(len(guests), happiness)
	fmt.Println(total)
}

