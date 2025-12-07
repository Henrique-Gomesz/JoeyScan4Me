package runner

import "sync"

func StartScan(opt *Options) {
	RunSubfinder(opt)
	RunHttpx(opt)

	var wg sync.WaitGroup

	// in order to speed up things we run katana and gowitness concurrently
	wg.Add(2)

	go func() {
		defer wg.Done()
		RunKatana(opt)
	}()

	go func() {
		defer wg.Done()
		RunGowitness(opt)
	}()

	wg.Wait()

	// start go witness server if --server flag is set
	if opt.Server {
		StartGoWitnessServer(opt)
	}
}
