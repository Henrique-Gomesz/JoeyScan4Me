package main

import (
	"JoeyScan4Me/internal/logging"
	"JoeyScan4Me/internal/orchestrator"
)

func main() {
	logging.PrintBanner()

	opt := orchestrator.ParseOptions()

	orchestrator.CheckToolSetup(opt)

	orchestrator.StartScan(opt)
}
