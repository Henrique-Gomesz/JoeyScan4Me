package config

type ToolsConfig struct {
	Tools []Tool `json:"tools"`
}

type Tool struct {
	Name           string `json:"name"`
	Command        string `json:"command"`
	InstallCommand string `json:"install_command"`
}
