package system

import (
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
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func init() {

}
func ImportSMS() ([]SMSData, error) {

	bytesRead, err := ioutil.ReadFile("sms.data")
	if err != nil {
		return []SMSData{}, err
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	SMSdata := make([]SMSData, 0)
	for _, v := range lines {
		sms, ok := parseLine(v)
		if ok {
			SMSdata = append(SMSdata, sms)
		}
	}
	return SMSdata, nil

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
