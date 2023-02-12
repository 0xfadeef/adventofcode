package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}
var origin = point{0,0}
const with_robot = true

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

	var santa, robot point = origin, origin
	var pos *point = &santa  // santa moves first

	visited := make(map[point]bool)
	visited[origin] = true

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
			switch buf[i] {
			case '>': pos.x++
			case '<': pos.x--
			case '^': pos.y++
			case 'v': pos.y--
			default: continue
			}
			if _, ok := visited[*pos]; !ok {
				visited[*pos] = true
			}
			if with_robot {
				if pos == &santa {
					pos = &robot
				} else {
					pos = &santa
				}
			}
		}
	}
	fmt.Println(len(visited))
}
