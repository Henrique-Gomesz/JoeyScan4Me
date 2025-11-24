package orchestrator

import (
	"sync"
)

func StarScan(opt *Options) {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		RunSubfinder(opt)
	}()

	wg.Wait()
}
