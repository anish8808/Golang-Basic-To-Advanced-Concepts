package main

import (
	"fmt"
	"sync"
	"time"
)

func count(thing string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count("fish")
		wg.Done()
	}()
	wg.Wait()
}
