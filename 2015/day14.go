package main

import (
	"bufio"
	"fmt"
	"os"
)


const duration = 2503

type Reindeer struct {
	name string
	speed, stamina, cooldown int
}
type State struct {
	points, distance int
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

	const pattern = "%s can fly %d km/s for %d seconds, but then must rest for %d seconds."
	scanner := bufio.NewScanner(f)

	var reindeers = make(map[Reindeer]*State)
	var max_distance = 0

	for scanner.Scan() {
		line := scanner.Text()
		r := Reindeer{}
		fmt.Sscanf(line, pattern, &r.name, &r.speed, &r.stamina, &r.cooldown)
		reindeers[r] = new(State)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	for r, _ := range reindeers {
		cycle := r.stamina + r.cooldown
		n_cycles, time_left := duration / cycle, duration % cycle
		d := (n_cycles * r.stamina + min(time_left, r.stamina)) * r.speed
		max_distance = max(max_distance, d)
	}
	fmt.Println(max_distance)


	for t := 0; t < duration; t++ {
		max_distance = 0
		for r, state := range reindeers {
			if t % (r.stamina + r.cooldown) < r.stamina {
				state.distance += r.speed
			}
			max_distance = max(max_distance, state.distance)
		}
		for _, state := range reindeers {
			if state.distance == max_distance {
				state.points += 1
			}
		}
	}

	max_points := 0
	for _, state := range reindeers {
		max_points = max(max_points, state.points)
	}
	fmt.Println(max_points)
}

