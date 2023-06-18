package main

import (
	"bufio"
	"fmt"
	"os"
)

type Edge struct {
	city1 string
	city2 string
}

var distances = make(map[Edge]uint)
var dummy string
var cities = make([]string, 0, 32)


func add_city(city string) bool {
	for _, c := range cities {
		if c == city {
			return false
		}
	}
	cities = append(cities, city)
	return true
}

func minmax_path_cost(start string, mask uint) (uint, uint) {
	var last = false
	var maxcost uint = 0
	var mincost uint = ^maxcost

	if mask & (mask - 1) == 0 {
		last = true
	}
	for k, city := range cities {
		ptr := uint(1 << k)

		if mask & ptr == 0 {
			continue
		}
		/* For dummy edges city_dist will be set to 0,
		   since they are not in the distance map.
		   Luckily, that is exactly what dummy egde
		   weight is supposed to be. */
		city_dist, ok := distances[Edge{start, city}]

		if !ok && start != dummy {
			continue
		}
		if last {
			mincost = city_dist
			maxcost = mincost
			break
		}
		locost, hicost := minmax_path_cost(city, mask &^ ptr)
		locost += city_dist; hicost += city_dist

		if locost < mincost {
			mincost = locost
		}
		if hicost > maxcost {
			maxcost = hicost
		}
	}
	return mincost, maxcost
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

	var city1, city2 string
	var dist uint
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscanf(line, "%s to %s = %d", &city1, &city2, &dist)

		distances[Edge{city1, city2}] = dist
		distances[Edge{city2, city1}] = dist

		add_city(city1)
		add_city(city2)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	n := len(cities)
	if n < 2 {
		fmt.Printf("Too few cities to proceed: %v\n", cities)
		return
	}

	mask := uint(1 << n) - 1
	mincost, maxcost := minmax_path_cost(dummy, mask)

	fmt.Println(mincost, maxcost)
}

