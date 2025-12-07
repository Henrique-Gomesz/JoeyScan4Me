package runner

import (
	"JoeyScan4Me/internal/logging"
	"fmt"
	"math"
	"path/filepath"

	"github.com/projectdiscovery/katana/pkg/engine/standard"
	katanaTypes "github.com/projectdiscovery/katana/pkg/types"
)

func RunKatana(opt *Options) {
	httpxOutputPath := filepath.Join(GetOutputFilePath(opt.Workdir, opt.Domain), HttpxOutputFile)
	katanaOutputPath := filepath.Join(GetOutputFilePath(opt.Workdir, opt.Domain), KatanaOutputFile)

	urls, err := ReadFileLines(httpxOutputPath)
	if err != nil {
		logging.LogError("Failed to read httpx output file", err)
	}

	if len(urls) == 0 {
		logging.LogInfo("No URLs found to crawl")
		return
	}

	file, err := CreateOutputFile(katanaOutputPath)
	if err != nil {
		logging.LogError("Failed to create katana output file", err)
	}
	defer file.Close()

	logging.LogInfo("Starting crawling with Katana")
	logging.LogInfo(fmt.Sprintf("Crawling %d URLs", len(urls)))

	katanaOpts := &katanaTypes.Options{
		MaxDepth:               3,
		BodyReadSize:           math.MaxInt,
		Timeout:                10,
		Concurrency:            20,
		Parallelism:            50,
		Retries:                3,
		RateLimit:              150,
		Strategy:               "depth-first",
		ScrapeJSResponses:      true,
		ScrapeJSLuiceResponses: true,
		OutputFile: katanaOutputPath,
	}

	crawlerOptions, err := katanaTypes.NewCrawlerOptions(katanaOpts)
	if err != nil {
		logging.LogError("Failed to create crawler options", err)
	}
	defer crawlerOptions.Close()

	crawler, err := standard.New(crawlerOptions)
	if err != nil {
		logging.LogError("Failed to create katana crawler", err)
	}
	defer crawler.Close()

	for _, url := range urls {
		if url == "" {
			continue
		}

		err = crawler.Crawl(url)

		if err != nil {
			logging.LogError(fmt.Sprintf("Could not crawl %s", url), err)
		}
	}

	logging.LogSuccess("Results saved to " + katanaOutputPath)
}
