package orchestrator

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

	for _, tool := range toolsConfig.Tools {
		logging.LogInfo("Checking tools")
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

func RunSubFinder(opt *Options) {
	logging.LogInfo("Running Subfinder")
	exec.Command("subfinder", "-d", opt.Domain, "--all", "--output", "subfinder.txt")
}
