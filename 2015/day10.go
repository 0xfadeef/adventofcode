package main

import (
	"fmt"
	"os"
	. "github.com/0xfadeef/look_and_say_elements"
)

const repeats1 = 40
const repeats2 = 50

type Chain []int

func (chain Chain) print_total_length() {
	result := 0
	for _, n := range chain {
		result += len(Elements[n].Sequence)
	}
	fmt.Println(result)
}

func (chain Chain) next_iteration() Chain {
	l := len(chain)
	for _, n := range chain {
		chain = append(chain, Elements[n].Products...)
	}
	return chain[l:]
}

func (chain Chain) evolve(repeats int) Chain {
	for i := 0; i < repeats; i++ {
		chain = chain.next_iteration()
	}
	return chain
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("missing input!\n")
		return
	}
	input := os.Args[1]
	var chain Chain

	for _, elem := range Elements {
		if input == elem.Sequence {
			chain = Chain([]int{elem.Number})
			break
		}
	}
	if chain == nil {
		// too lazy to implement compound input case
		fmt.Println("Input must be a simple element!")
		return
	}

	chain = chain.evolve(repeats1)
	chain.print_total_length()

	chain = chain.evolve(repeats2-repeats1)
	chain.print_total_length()
}

