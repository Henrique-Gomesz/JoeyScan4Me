package logging

import "github.com/fatih/color"

func PrintBanner() {
	color.Magenta(`
		JoeyScan4Me - Recon toolkit

               |\__/,|   (\                       
             _.|o o  |_   ) )                     
           -(((---(((--------  
			by: Henrique-Gomesz                  							  
`)
}

func LogError(message string, err error) {
	color.Red("[x] %s: %v", message, err)
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
