package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const exclude_red = true

type (
	jsonArray  = []interface{}
	jsonObject = map[string]interface{}
	jsonNumber = float64
)

func string_value_red(v any) bool {
	s, ok := v.(string)
	return ok && s == "red"
}

func traverse(object any) int {
	result := 0

	switch o := object.(type) {
	case jsonArray:
		for _, v := range o {
			result += traverse(v)
		}
	case jsonObject:
		for _, v := range o {
			if exclude_red && string_value_red(v) {
				return 0
			}
			result += traverse(v)
		}
	case jsonNumber:
		return int(o)
	}
	return result
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing input file")
		return
	}
	input := os.Args[1]
	data, err := os.ReadFile(input)

	if err != nil {
		fmt.Println(err)
		return
	}

	var object interface{}
	err = json.Unmarshal(data, &object)

	if err != nil {
		fmt.Println(err)
		return
	}
	result := traverse(object)
	fmt.Println(result)
}

