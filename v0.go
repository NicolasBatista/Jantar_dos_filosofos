package main

import (
	"fmt"
	"time"
)

func Philosopher(id, iteration int) {
	time_to_think := 2
	time_to_eat := 3

	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" T%d\n", iteration)
	time.Sleep(time.Duration(time_to_think) * time.Millisecond)

	for s := 1; s <= id*6; s++ {
		fmt.Printf(" ")
	}
	fmt.Printf(" E%d\n", iteration)
	time.Sleep(time.Duration(time_to_eat) * time.Millisecond)
}

func RodarV0() {
	philosophers := 5
	rounds := 20

	fmt.Println("\n[P1] [P2] [P3] [P4] [P5]\n")

	start := time.Now()

	for i := 1; i <= rounds; i++ {
		for j := 0; j < philosophers; j++ {
			Philosopher(j, i)
		}
	}
	elapsed := time.Since(start)

	fmt.Printf("\nTempo sequencial: %s\n\n", elapsed)
}
