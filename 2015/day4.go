package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("missing input!\n")
		return
	}
	secret := os.Args[1]
	secret_length := len(secret)
	buffer_length := secret_length + 10  // max digits in unit32

	buffer := make([]byte, secret_length, buffer_length)
	copy(buffer, secret)

	var number uint32 = 1
	var found5 = false
	for {
		data := fmt.Append(buffer, number)
		hash := md5.Sum(data)
		hexhash := fmt.Sprintf("%x", hash)

		if hexhash[:5] == "00000" {
			if !found5 {
				fmt.Println(number)
				found5 = true
			}
			if hexhash[5] == '0' {
				fmt.Println(number)
				break
			}
		}
		number += 1
	}
}
