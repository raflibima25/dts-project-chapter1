package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	var mtx sync.Mutex
	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go bisa(i, &wg)
	}

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go coba(i, &wg)
	}

	wg.Wait()

	fmt.Println(strings.Repeat("-", 20))

	// with mutex
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		mtx.Lock()
		go bisaMtx(i, &mtx, &wg)
		mtx.Unlock()
	}

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		mtx.Lock()
		go cobaMtx(i, &mtx, &wg)
		mtx.Unlock()
	}

	wg.Wait()
}

func bisa(index int, wg *sync.WaitGroup) {
	fmt.Printf("[bisa1 bisa2 bisa3] %d \n", index)
	defer wg.Done()
}

func coba(index int, wg *sync.WaitGroup) {
	fmt.Printf("[coba1 coba2 coba3] %d \n", index)
	defer wg.Done()
}

func bisaMtx(index int, mtx *sync.Mutex, wg *sync.WaitGroup) {
	fmt.Printf("[bisa1 bisa2 bisa3] %d \n", index)
	defer wg.Done()
}

func cobaMtx(index int, mtx *sync.Mutex, wg *sync.WaitGroup) {
	fmt.Printf("[coba1 coba2 coba3] %d \n", index)
	defer wg.Done()
}
