package runner

import (
	"JoeyScan4Me/internal/logging"
	"os"
	"strings"

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
	flagSet.StringVar(&opt.Workdir, "w", "./", "working directory for output files, defaults to current directory")

	if err := flagSet.Parse(); err != nil {
		logging.LogError("Error parsing flags:", err)
		os.Exit(1)
	}

	validateDomain(opt)

	return opt
}

func validateDomain(opt *Options) {
	if opt.Domain == "" {
		logging.LogError("Domain is required. Use -d flag to specify a domain (e.g., -d example.com)", nil)
		os.Exit(1)
	}

	opt.Domain = strings.TrimPrefix(opt.Domain, "https://")
	opt.Domain = strings.TrimPrefix(opt.Domain, "http://")

	opt.Domain = strings.TrimSuffix(opt.Domain, "/")
}
