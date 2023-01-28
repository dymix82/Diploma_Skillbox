package sms

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	CountrySms = iota
	BandwidthSms
	ResponseTimeSms
	ProviderSms
)

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"provider"`
	ResponseTime string `json:"bandwidth"`
	Provider     string `json:"response_time"`
}

func init() {
	//	Providers := [3]string{"Topolo", "Rond", "Kildy"}
}
func ImportSMS() {
	bytesRead, _ := ioutil.ReadFile("sms.data")
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	SMSdata := make([]SMSData, 0)
	for _, v := range lines {
		sms, ok := parseLine(v)
		if ok {
			SMSdata = append(SMSdata, sms)
		}
	}
	fmt.Println("SMS: ", SMSdata)
	//for i, _ := range SMSdata {
	//	fmt.Println(SMSdata[i])
	//}

}
func parseLine(line string) (SMSData, bool) {
	sms := strings.Split(line, ";")

	switch {
	case len(sms) != 4:
		fallthrough
	case !isCountryOK(sms[CountrySms]):
		fallthrough
	case !isProviderSMSOK(sms[ProviderSms]):
		return SMSData{}, false
	}
	return SMSData{
		Country:      sms[CountrySms],
		Bandwidth:    sms[BandwidthSms],
		ResponseTime: sms[ResponseTimeSms],
		Provider:     sms[ProviderSms],
	}, true
}
