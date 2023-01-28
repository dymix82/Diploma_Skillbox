package main

import "main/pkg/sms"

func main() {
	sms.ImportSMS()
	sms.ImportMMS()
	sms.ImportVoice()
}
