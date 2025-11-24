package orchestrator

import (
	"JoeyScan4Me/internal/logging"
	"os"

	"github.com/projectdiscovery/goflags"
)

type Options struct {
	Domain  string
	Workdir string
}

func ParseOptions() *Options {
	opt := &Options{}
	flagSet := goflags.NewFlagSet()

	flagSet.SetDescription("JoeyScan4Me toolkit")

	flagSet.StringVar(&opt.Domain, "d", "", "domain to scan (e.g. example.com)")
	flagSet.StringVar(&opt.Workdir, "w", "./", "working directory for output files")

	if err := flagSet.Parse(); err != nil {
		logging.LogError("Error parsing flags:", err)
		os.Exit(1)
	}

	return opt
}
