package runner

import (
	"JoeyScan4Me/internal/logging"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	httpRunner "github.com/projectdiscovery/httpx/runner"
)

func RunHttpx(opt *Options) {
	outputPath := filepath.Join(GetOutputFilePath(opt.Workdir, "subdomains", opt.Domain), HttpxOutputFile)
	techOutputPath := filepath.Join(GetOutputFilePath(opt.Workdir, "subdomains", opt.Domain), HttpxTechOutputFile)
	subfinderFile := filepath.Join(GetOutputFilePath(opt.Workdir, "subdomains", opt.Domain), SubfinderOutputFile)

	file, err := CreateOutputFile(outputPath)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()

	techFile, err := CreateOutputFile(techOutputPath)
	if err != nil {
		log.Fatalf("Failed to create tech output file: %v", err)
	}
	defer techFile.Close()

	httpxOpts := &httpRunner.Options{
		Methods:         "GET",
		FollowRedirects: true,
		TechDetect:      true,
		RandomAgent:     true,
		InputFile:       subfinderFile,
		OnResult: func(r httpRunner.Result) {
			if r.Err != nil {
				logging.LogError("HTTPX error:", r.Err)
				return
			}

			fmt.Fprintf(file, "%s\n", r.URL)

			if len(r.Technologies) > 0 {
				techList := strings.Join(r.Technologies, ", ")
				fmt.Fprintf(techFile, "%s [%s]\n", r.URL, techList)
			} else {
				fmt.Fprintf(techFile, "%s [no technologies detected]\n", r.URL)
			}
		},
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
