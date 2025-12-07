package runner

import (
	"os"
	"path/filepath"

	"JoeyScan4Me/internal/logging"
)

var SubfinderOutputFile = "subdomains.txt"
var HttpxOutputFile = "up_subdomains.txt"
var KatanaOutputFile = "crawling_results.txt"

func GetOutputFilePath(workdir, tool, domain string) string {
	return filepath.Join(workdir, "output", domain)
}

func CreateOutputFile(filePath string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		logging.LogError("Failed to create output directory", err)
		return nil, err
	}

	file, err := os.Create(filePath)
	if err != nil {
		logging.LogError("Failed to create output file", err)
		return nil, err
	}

	return file, nil
}

