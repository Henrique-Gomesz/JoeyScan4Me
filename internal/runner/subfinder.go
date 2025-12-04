package runner

import (
	"JoeyScan4Me/internal/logging"
	"bytes"
	"context"
	"io"
	"log"
	"path/filepath"
	subfinderRunner "github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

func RunSubfinder(opt *Options) {
	filePath := filepath.Join(GetOutputFilePath(opt.Workdir, "subdomains", opt.Domain), SubfinderOutputFile)

	subfinderOpts := &subfinderRunner.Options{
		Threads:            10, 
		Timeout:            30, 
		MaxEnumerationTime: 10, 
		OutputDirectory: filePath,
	}
	
	subfinder, err := subfinderRunner.NewRunner(subfinderOpts)
	if err != nil {
		logging.LogError("failed to create subfinder runner: %v", err)
	}

	output := &bytes.Buffer{}
	var sourceMap map[string]map[string]struct{}

	if sourceMap, err = subfinder.EnumerateSingleDomainWithCtx(context.Background(), opt.Domain, []io.Writer{output}); err != nil {
		log.Fatalf("failed to enumerate single domain: %v", err)
	}

	log.Printf("sourceMap dump: %+v", sourceMap)
	for source, domains := range sourceMap {
		log.Printf("Source: %s", source)
		for d := range domains {
			log.Printf("  - %s", d)
		}
	}

	logging.LogInfo("Searching for subdomains")
	logging.LogInfo("Saving subdomains results to " + filePath)
}