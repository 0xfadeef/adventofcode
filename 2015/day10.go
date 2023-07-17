package main

import (
	"fmt"
	"os"
	. "github.com/0xfadeef/look_and_say_elements"
)

const repeats = 40

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("missing input!\n")
		return
	}
	input := os.Args[1]
	var chain []int

	for _, elem := range Elements {
		if input == elem.Sequence {
			chain = append(chain, elem.Number)
			break
		}
	}
	if chain == nil {
		// too lazy to implement compound input case
		fmt.Println("Input must be a simple element!")
		return
	}

	for i := 0; i < repeats; i++ {
		l := len(chain)
		for _, n := range chain {
			chain = append(chain, Elements[n].Products...)
		}
		chain = chain[l:]
	}

	var result = 0
	for _, n := range chain {
		result += len(Elements[n].Sequence)
	}
	fmt.Println(result)
}

