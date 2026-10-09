package main

import (
	"fmt"
	"sync"
	"time"
)

func RodarV1() {
	philosophers := 5
	rounds := 20

	fmt.Println("\n[P1] [P2] [P3] [P4] [P5]\n")

	var forks [5]sync.Mutex
	var wg sync.WaitGroup

	wg.Add(philosophers)

	for i := 0; i < philosophers; i++ {
		go func(id int) {
			esq := id
			dir := (id + 1) % 5

			for j := 1; j <= rounds; j++ {
				espacos := ""
				for s := 1; s <= id*6; s++ {
					espacos += " "
				}
				fmt.Printf("%s T%d\n", espacos, j)
				time.Sleep(2 * time.Millisecond)

				forks[esq].Lock()
				// delay para causar o deadlock
				time.Sleep(10 * time.Millisecond)
				forks[dir].Lock()

				espacos = ""
				for s := 1; s <= id*6; s++ {
					espacos += " "
				}
				fmt.Printf("%s E%d\n", espacos, j)
				time.Sleep(3 * time.Millisecond)

				forks[dir].Unlock()
				forks[esq].Unlock()
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
}
