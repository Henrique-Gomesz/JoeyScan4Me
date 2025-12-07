package main

import (
	"github.com/Henrique-Gomesz/JoeyScan4Me/pkg/logging"
	"github.com/Henrique-Gomesz/JoeyScan4Me/pkg/runner"
)

var Version string = "1.0.0"

func main() {
	logging.PrintBanner(Version)

	opt := runner.ParseOptions()

	runner.StartScan(opt)
}
