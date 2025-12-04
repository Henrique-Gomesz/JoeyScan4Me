package main

import (
	"JoeyScan4Me/internal/logging"
	"JoeyScan4Me/internal/runner"
)
	
func main() {
	logging.PrintBanner()

	opt := runner.ParseOptions()

	//runner.CheckToolSetup(opt)

	runner.StartScan(opt)
}
