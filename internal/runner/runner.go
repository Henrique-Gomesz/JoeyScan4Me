package runner

func StartScan(opt *Options) {
	

	//var wg sync.WaitGroup
	
	RunSubfinder(opt)
	// RunHttpx(opt)

	// go func() {
	// 	defer wg.Done()
	// 	RunKatana(opt)
	// }()

	// wg.Wait()
}
