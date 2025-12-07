package runner

import (
	"JoeyScan4Me/internal/logging"
	"context"
	"log"
	"os"
	"path/filepath"

	subfinderRunner "github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

func RunSubfinder(opt *Options) {
	filePath := filepath.Join(GetOutputFilePath(opt.Workdir, "subdomains", opt.Domain), SubfinderOutputFile)
	file, err := CreateOutputFile(filePath)

	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
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
		log.Fatalf("Failed to create subfinder runner: %v", err)
	}

	if err = subfinder.RunEnumerationWithCtx(context.Background()); err != nil {
		logging.LogError("Failed to enumerate subdomains", err)

		// end process execution because subfinder is essential to run the next tools;
		os.Exit(1)
	}
}
