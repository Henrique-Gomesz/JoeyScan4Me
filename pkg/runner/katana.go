package runner

import (
	"fmt"
	"math"
	"net/url"
	"os"
	"path/filepath"
	"sync"

	"github.com/henrique-gomesz/joeyscan4me/pkg/logging"

	"github.com/projectdiscovery/katana/pkg/engine/standard"
	katanaOutput "github.com/projectdiscovery/katana/pkg/output"
	katanaTypes "github.com/projectdiscovery/katana/pkg/types"
)

var KatanaCrawlingDir = "crawling"

func RunKatana(opt *Options) error {
	httpxOutputPath := filepath.Join(GetOutputFilePath(opt.Workdir, opt.Domain), HttpxOutputFile)
	crawlingDir := filepath.Join(GetOutputFilePath(opt.Workdir, opt.Domain), KatanaCrawlingDir)

	if opt.Resume && dirNonEmpty(crawlingDir) {
		logging.LogInfo("Skipping katana — crawling directory already exists: " + crawlingDir)
		return nil
	}

	urls, err := ReadFileLines(httpxOutputPath)
	if err != nil {
		return fmt.Errorf("failed to read httpx output file: %w", err)
	}
	urls = NormalizeAndDedupeLines(urls)

	if len(urls) == 0 {
		logging.LogInfo("No URLs found to crawl")
		return nil
	}

	if err := os.MkdirAll(crawlingDir, 0755); err != nil {
		return fmt.Errorf("failed to create crawling directory: %w", err)
	}

	logging.LogInfo("Starting crawling with Katana")
	logging.LogInfo(fmt.Sprintf("Crawling %d URLs", len(urls)))

	var mu sync.Mutex
	openFiles := make(map[string]*os.File)

	defer func() {
		for _, f := range openFiles {
			f.Close()
		}
	}()

	katanaOpts := &katanaTypes.Options{
		URLs:                   urls,
		MaxDepth:               opt.KatanaDepth,
		BodyReadSize:           math.MaxInt,
		Timeout:                opt.KatanaTimeout,
		Concurrency:            opt.KatanaConcurrency,
		Parallelism:            opt.KatanaParallelism,
		Retries:                3,
		RateLimit:              opt.KatanaRateLimit,
		Strategy:               "depth-first",
		ScrapeJSResponses:      true,
		ScrapeJSLuiceResponses: true,
		ExtensionFilter:        []string{"css"},
		OnResult: func(result katanaOutput.Result) {
			discovered := result.Request.URL
			if discovered == "" {
				return
			}

			host := hostFromURL(discovered)
			if host == "" {
				host = "unknown"
			}

			mu.Lock()
			defer mu.Unlock()

			f, ok := openFiles[host]
			if !ok {
				hostFile := filepath.Join(crawlingDir, host+".txt")
				var ferr error
				f, ferr = os.OpenFile(hostFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if ferr != nil {
					logging.LogError("Failed to open crawl file for "+host, ferr)
					return
				}
				openFiles[host] = f
			}

			fmt.Fprintf(f, "%s\n", discovered)
		},
	}

	crawlerOptions, err := katanaTypes.NewCrawlerOptions(katanaOpts)
	if err != nil {
		return fmt.Errorf("failed to create katana crawler options: %w", err)
	}
	defer crawlerOptions.Close()

	crawler, err := standard.New(crawlerOptions)
	if err != nil {
		return fmt.Errorf("failed to create katana crawler: %w", err)
	}
	defer crawler.Close()

	failed := 0
	for _, targetURL := range urls {
		if err := crawler.Crawl(targetURL); err != nil {
			failed++
			logging.LogError(fmt.Sprintf("Could not crawl %s", targetURL), err)
		}
	}

	if failed > 0 {
		return fmt.Errorf("katana failed to crawl %d/%d URLs", failed, len(urls))
	}

	logging.LogSuccess(fmt.Sprintf("Crawling results saved to %s/", crawlingDir))
	return nil
}

// hostFromURL extracts the hostname (without port) from a raw URL string.
func hostFromURL(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil || u.Host == "" {
		return ""
	}
	host := u.Hostname()
	return host
}

// dirNonEmpty returns true when the directory exists and contains at least one file.
func dirNonEmpty(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	for _, e := range entries {
		if !e.IsDir() {
			return true
		}
	}
	return false
}
