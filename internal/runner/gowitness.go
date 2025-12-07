package runner

import (
	"JoeyScan4Me/internal/logging"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/sensepost/gowitness/pkg/log"
	goWitnessRunner "github.com/sensepost/gowitness/pkg/runner"
	driver "github.com/sensepost/gowitness/pkg/runner/drivers"
	"github.com/sensepost/gowitness/pkg/writers"
)

// thx Claude Sonnet
func RunGowitness(opt *Options) {
	httpxOutputFile := filepath.Join(GetOutputFilePath(opt.Workdir, opt.Domain), HttpxOutputFile)
	screenshotsPath := filepath.Join(GetOutputFilePath(opt.Workdir, opt.Domain), "screenshots")

	if _, err := os.Stat(httpxOutputFile); os.IsNotExist(err) {
		logging.LogError("HTTPX output file not found. Run HTTPX first.", err)
		return
	}

	urls, err := ReadFileLines(httpxOutputFile)
	if err != nil {
		logging.LogError("Failed to read URLs from httpx output:", err)
		return
	}

	if len(urls) == 0 {
		logging.LogInfo("No URLs found to screenshot")
		return
	}

	logging.LogInfo("Running gowitness to capture screenshots from up domains")

	// Configure gowitness options
	options := goWitnessRunner.NewDefaultOptions()
	options.Scan.ScreenshotPath = screenshotsPath
	options.Scan.ScreenshotFormat = "jpeg"
	options.Scan.ScreenshotJpegQuality = 80
	options.Scan.Threads = 5
	options.Scan.Timeout = 30
	options.Scan.Delay = 3
	options.Logging.LogScanErrors = true

	// Create slog logger
	logger := slog.New(log.Logger)

	// Initialize chromedp driver
	gwDriver, err := driver.NewChromedp(logger, *options)
	if err != nil {
		logging.LogError("Failed to create gowitness driver:", err)
		return
	}
	defer gwDriver.Close()

	// Create runner with no writers (only screenshots)
	gwRunner, err := goWitnessRunner.NewRunner(logger, gwDriver, *options, []writers.Writer{})
	if err != nil {
		logging.LogError("Failed to create gowitness runner:", err)
		return
	}
	defer gwRunner.Close()

	// Feed URLs to the runner
	go func() {
		for _, url := range urls {
			gwRunner.Targets <- url
		}
		close(gwRunner.Targets)
	}()

	// Run the screenshot capture
	gwRunner.Run()

	logging.LogInfo("Gowitness completed. Screenshots saved in: " + screenshotsPath)
}
