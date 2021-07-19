package main

import (
	"go-design-patterns/singleton"
	"sync"
)

var logger = &singleton.InfoLogger{}
func main () {
	var wg sync.WaitGroup
	for i:=0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			logger.Pintf("log: %v", i)

		}(i)
	}

	wg.Wait()
}