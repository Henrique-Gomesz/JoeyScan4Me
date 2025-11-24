package orchestrator

import (
	"sync"
)

func StarScan(opt *Options) {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		RunSubFinder(opt)
	}()

	wg.Wait()
}
