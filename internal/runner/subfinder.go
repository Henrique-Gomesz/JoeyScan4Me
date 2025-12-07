package runner

import (
	"JoeyScan4Me/internal/logging"
	"context"
	"os"
	"path/filepath"

	subfinderRunner "github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

func RunSubfinder(opt *Options) {
	filePath := filepath.Join(GetOutputFilePath(opt.Workdir, opt.Domain), SubfinderOutputFile)
	file, err := CreateOutputFile(filePath)

	if err != nil {
		logging.LogError("Failed to create output file:", err)
	}

	defer file.Close()

	subfinderOpts := &subfinderRunner.Options{
		Threads:            10,
		Timeout:            30,
		MaxEnumerationTime: 10,
		All:                true,
		Domain:             []string{opt.Domain},
		OutputFile:         filePath,
		OutputDirectory:    filepath.Dir(filePath),
		Output:             file,
	}

	subfinder, err := subfinderRunner.NewRunner(subfinderOpts)
	if err != nil {
		logging.LogError("Failed to create subfinder runner", err)
	}

	logging.LogInfo("Running subfinder")
	if err = subfinder.RunEnumerationWithCtx(context.Background()); err != nil {
		logging.LogError("Failed to enumerate subdomains", err)

		// end process execution because subfinder is essential to run the next tools;
		os.Exit(1)
	}
}
