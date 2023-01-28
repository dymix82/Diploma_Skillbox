package main

import "main/pkg/system"

func main() {
	system.ImportSMS()
	system.ImportMMS()
	system.ImportVoice()
	system.ImportEmail()
}
