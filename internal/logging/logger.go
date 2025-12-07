package logging

import "github.com/fatih/color"

func PrintBanner() {
	color.Magenta(`
JoeyScan4Me - Simple and helpful recon toolkit

    |\__/,|   ('\
  _.|o o  |_   ) )
-(((---(((--------
by: Henrique-Gomesz              							  
`)
}

func LogError(message string, err error) {
	if err != nil {
		color.Red("[x] %s: %v", message, err)
	} else {
		color.Red("[x] %s", message)
	}
}

func LogInfo(message string) {
	color.Cyan("[i] %s", message)
}

func LogSuccess(message string) {
	color.Green("[✓] %s", message)
}

func LogText(message string) {
	color.White("%s", message)
}
