package main

import (
	"JoeyScan4Me/internal/logging"
	"JoeyScan4Me/internal/runner"
)
	
func main() {
	logging.PrintBanner()

	opt := runner.ParseOptions()

	runner.StartScan(opt)
}
