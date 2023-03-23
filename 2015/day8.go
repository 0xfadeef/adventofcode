package main

import (
	"bufio"
	"fmt"
	"os"
)

func isHex(ch byte) bool {
	return ('0' <= ch && ch <= '9') || ('a' <= ch && ch <= 'f') || ('A' <= ch && ch <= 'F')
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

	var diff, hexed = 0, 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		size := len(line)

		if size < 2 || line[0] != '"' || line[size-1] != '"' {
			fmt.Println("invalid input: unquoted line")
			return
		}
		diff += 2

		for i := 1; i < size-1; i++ {
			if line[i] != '\\' {
				continue
			} else if i < size-2 {
				// inspect next character
				i += 1
			} else {
				fmt.Println("invalid input: closing quote is escaped")
				return
			}
			switch line[i] {
			case 'x':
				if i < size-3 && isHex(line[i+1]) && isHex(line[i+2]) {
					i += 2
					hexed += 1
				} else {
					fmt.Println("invalid input: bad hex value")
					return
				}
			case '\\', '"':
				diff += 1
			default:
				fmt.Println("invalid input: unknown escape sequence")
				return
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Println(diff + 3*hexed, 2*diff + hexed)
	}
}

