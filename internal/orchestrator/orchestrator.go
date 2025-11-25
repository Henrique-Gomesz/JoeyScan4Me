package orchestrator

import (
	"sync"
)

func StartScan(opt *Options) {
	var wg sync.WaitGroup

	RunSubfinder(opt)
	RunHttpx(opt)

	wg.Add(1)

	go func() {
		defer wg.Done()
		RunKatana(opt)
	}()

	wg.Wait()
}
