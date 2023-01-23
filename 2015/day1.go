package main

import (
	"fmt"
	"os"
)

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

	buf := make([]byte, 1024)
	floor := 0

	const basement = -1
	basement_at := 0
	found := false

	for {
		n, err := f.Read(buf)

		if !(n > 0) {
			// end of file reached
			break
		} else if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		for i := 0; i < n; i++ {
			if buf[i] == '(' {
				floor++
			} else if buf[i] == ')' {
				floor--
			} else {
				continue
			}
			if !found {
				basement_at++
				if floor == basement {
					found = true
				}
			}
		}
	}
	fmt.Println(floor, basement_at)
}

