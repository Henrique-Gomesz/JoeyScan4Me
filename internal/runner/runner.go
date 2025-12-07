package runner

func StartScan(opt *Options) {

	RunSubfinder(opt)
	RunHttpx(opt)
	RunKatana(opt)
}
