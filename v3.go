package main

import (
	"fmt"
	"sync"
	"time"
)

var totalEsperaV3 int64
var maiorEsperaV3 int64
var maxComendoV3 int
var comendoAgoraV3 int
var mutexMetricasV3 sync.Mutex

func RodarV3() {
	philosophers := 5
	rounds := 20

	fmt.Println("\n[P1] [P2] [P3] [P4] [P5]\n")
	inicio := time.Now()

	var forks [5]sync.Mutex
	var wg sync.WaitGroup

	wg.Add(philosophers)

	for i := 0; i < philosophers; i++ {
		go func(id int) {
			esq := id
			dir := (id + 1) % 5

			primeiro := esq
			segundo := dir
			if esq > dir {
				primeiro = dir
				segundo = esq
			}

			for j := 1; j <= rounds; j++ {
				espacos := ""
				for s := 1; s <= id*6; s++ {
					espacos += " "
				}
				fmt.Printf("%s T%d\n", espacos, j)
				time.Sleep(2 * time.Millisecond)

				esperaInicio := time.Now()

				forks[primeiro].Lock()
				forks[segundo].Lock()

				esperou := time.Since(esperaInicio).Milliseconds()

				mutexMetricasV3.Lock()
				totalEsperaV3 += esperou
				if esperou > maiorEsperaV3 {
					maiorEsperaV3 = esperou
				}
				comendoAgoraV3++
				if comendoAgoraV3 > maxComendoV3 {
					maxComendoV3 = comendoAgoraV3
				}
				mutexMetricasV3.Unlock()

				espacos = ""
				for s := 1; s <= id*6; s++ {
					espacos += " "
				}
				fmt.Printf("%s E%d\n", espacos, j)
				time.Sleep(3 * time.Millisecond)

				mutexMetricasV3.Lock()
				comendoAgoraV3--
				mutexMetricasV3.Unlock()

				forks[segundo].Unlock()
				forks[primeiro].Unlock()
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
	tempoTotal := time.Since(inicio)

	fmt.Printf("\nTempo total V3: %s\n", tempoTotal)
	fmt.Printf("Tempo médio de espera: %d ms\n", totalEsperaV3/(int64(philosophers*rounds)))
	fmt.Printf("Maior tempo de espera: %d ms\n", maiorEsperaV3)
	fmt.Printf("Máximo comendo junto: %d\n\n", maxComendoV3)
}
