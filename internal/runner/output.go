package runner

import "path/filepath"

var SubfinderOutputFile = "subdomains.txt"
var HttpxOutputFile = "up_subdomains.txt"
var KatanaOutputFile = "crawling_results.txt"

func GetOutputFilePath(workdir, tool, domain string) string {
	return filepath.Join(workdir, "output", domain, tool)
}
