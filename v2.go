package main

import (
	"fmt"
	"sync"
	"time"
)

var totalEsperaV2 int64
var maiorEsperaV2 int64
var maxComendoV2 int
var comendoAgoraV2 int
var mutexMetricasV2 sync.Mutex

func RodarV2() {
	philosophers := 5
	rounds := 20

	fmt.Println("\n[P1] [P2] [P3] [P4] [P5]\n")
	inicio := time.Now()

	var mesa sync.Mutex
	var wg sync.WaitGroup

	wg.Add(philosophers)

	for i := 0; i < philosophers; i++ {
		go func(id int) {
			for j := 1; j <= rounds; j++ {
				espacos := ""
				for s := 1; s <= id*6; s++ {
					espacos += " "
				}
				fmt.Printf("%s T%d\n", espacos, j)
				time.Sleep(2 * time.Millisecond)

				esperaInicio := time.Now()

				mesa.Lock()

				esperou := time.Since(esperaInicio).Milliseconds()

				mutexMetricasV2.Lock()
				totalEsperaV2 += esperou
				if esperou > maiorEsperaV2 {
					maiorEsperaV2 = esperou
				}
				comendoAgoraV2++
				if comendoAgoraV2 > maxComendoV2 {
					maxComendoV2 = comendoAgoraV2
				}
				mutexMetricasV2.Unlock()

				espacos = ""
				for s := 1; s <= id*6; s++ {
					espacos += " "
				}
				fmt.Printf("%s E%d\n", espacos, j)
				time.Sleep(3 * time.Millisecond)

				mutexMetricasV2.Lock()
				comendoAgoraV2--
				mutexMetricasV2.Unlock()

				mesa.Unlock()
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
	tempoTotal := time.Since(inicio)

	fmt.Printf("\nTempo total V2: %s\n", tempoTotal)
	fmt.Printf("Tempo médio de espera: %d ms\n", totalEsperaV2/(int64(philosophers*rounds)))
	fmt.Printf("Maior tempo de espera: %d ms\n", maiorEsperaV2)
	fmt.Printf("Máximo comendo junto: %d\n\n", maxComendoV2)
}
