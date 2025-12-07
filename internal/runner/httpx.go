package runner

import (
	"JoeyScan4Me/internal/logging"
	"fmt"
	"log"
	"path/filepath"

	httpRunner "github.com/projectdiscovery/httpx/runner"
)

func RunHttpx(opt *Options) {
	outputPath := filepath.Join(GetOutputFilePath(opt.Workdir, "subdomains", opt.Domain), HttpxOutputFile)
	//subfinderFile := filepath.Join(GetOutputFilePath(opt.Workdir, "subdomains", opt.Domain), SubfinderOutputFile)
	file, err := CreateOutputFile(outputPath)

	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}

	defer file.Close()

	httpxOpts := &httpRunner.Options{
		Methods:         "GET",
		FollowRedirects: true,
		TechDetect:      true,
		RandomAgent:     true,
		OnResult: func(r httpRunner.Result) {
			if r.Err != nil {
				logging.LogError("HTTPX error:", r.Err)
				return
			}

			logging.LogInfo(fmt.Sprintf("Found alive subdomain: %s", r.URL))
		},
		Output: outputPath,
	}

	if err := httpxOpts.ValidateOptions(); err != nil {
		log.Fatal(err)
	}

	httpx, err := httpRunner.New(httpxOpts)
	if err != nil {
		logging.LogError("Failed to create httpx runner", err)
		log.Fatalf("Failed to create httpx runner: %v", err)
	}

	defer httpx.Close()

	httpx.RunEnumeration()
}
