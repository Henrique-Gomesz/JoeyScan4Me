package orchestrator

import "path/filepath"

func GetOutputFilePath(workdir, tool, fileName, domain string) string {
	return filepath.Join(workdir, "output", domain, tool, fileName)
}
