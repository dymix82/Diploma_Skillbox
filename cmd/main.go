package main

import (
	"main/pkg/handlers"
	"main/pkg/result"
	"main/pkg/system"
)

func main() {
	handlers.Router()
	system.ImportSMS()
	system.ImportMMS()
	system.ImportVoice()
	system.ImportEmail()
	system.ImportBilling()
	system.ImportSupport()
	system.ImportIncident()
	result.MakeSMSResult()
	result.MakeMMSResult()

}
