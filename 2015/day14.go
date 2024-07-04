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

	var reindeers = make([]Reindeer, 0, 32)
	var max_distance = 0

	for scanner.Scan() {
		line := scanner.Text()
		r := new(Reindeer)
		fmt.Sscanf(line, pattern, &r.name, &r.speed, &r.stamina, &r.cooldown)
		reindeers = append(reindeers, *r)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	for _, r := range reindeers {
		cycle := r.stamina + r.cooldown
		n_cycles, time_left := duration / cycle, duration % cycle
		d := (n_cycles * r.stamina + min(time_left, r.stamina)) * r.speed
		max_distance = max(max_distance, d)
	}
	fmt.Println(max_distance)
}

