package runner

import (
	"JoeyScan4Me/config"
	"JoeyScan4Me/internal/logging"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func CheckToolSetup(opt *Options) error {
	configPath := filepath.Join("..", "..", "config", "config.json")
	file, err := os.Open(configPath)
	if err != nil {
		logging.LogError("Error opening config file:", err)
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		logging.LogError("Error reading file:", err)
		os.Exit(1)
	}

	var toolsConfig config.ToolsConfig

	err = json.Unmarshal(buffer, &toolsConfig)
	if err != nil {
		logging.LogError("Error unmarshaling config JSON:", err)
		os.Exit(1)
	}

	logging.LogInfo("Checking tools")

	for _, tool := range toolsConfig.Tools {
		_, err := exec.LookPath(tool.Name)
		if err != nil {
			logging.LogInfo(fmt.Sprintf("Tool '%s' not found. Check if tool is correctly set on PATH or install it using: %s", tool.Name, tool.InstallCommand))
			os.Exit(1)
		} else {
			logging.LogSuccess(fmt.Sprintf("Tool '%s'", tool.Name))
		}
	}

	return nil
}

func RunHttpx(opt *Options) {
	logging.LogInfo("Checking alive subdomains")

	filePath := filepath.Join(GetOutputFilePath(opt.Workdir, "subdomains", opt.Domain), HttpxOutputFile)
	cmd := exec.Command("httpx", "-l", filepath.Join(GetOutputFilePath(opt.Workdir, "subdomains", opt.Domain), SubfinderOutputFile), "-o", filePath)

	cmd.Run()
	cmd.Wait()
	logging.LogInfo("Saving up subdomains results to " + filePath)
}

func RunKatana(opt *Options) {
	logging.LogInfo("Crawling alive subdomains")

	filePath := filepath.Join(GetOutputFilePath(opt.Workdir, "crawling", opt.Domain), KatanaOutputFile)

	cmd := exec.Command("katana", "-list", filepath.Join(GetOutputFilePath(opt.Workdir, "subdomains", opt.Domain), HttpxOutputFile), "-o", filePath)

	cmd.Run()
	cmd.Wait()
	logging.LogInfo("Saving crawling results to " + filePath)
}
