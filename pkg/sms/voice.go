package sms

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	CountryCall = iota
	BandwidthCall
	ResponseTimeCall
	ProviderCall
	ConnectionStabilityCall
	TTFBCall
	VoicePurityCall
	MedianOfCallsTimeCall
)

type VoiceData struct {
	Country             string
	Bandwidth           string
	ResponseTime        string
	Provider            string
	ConnectionStability string
	TTFB                string
	VoicePurity         string
	MedianOfCallsTime   string
}

func init() {
	//	Providers := [3]string{"Topolo", "Rond", "Kildy"}
}
func ImportVoice() {
	bytesRead, _ := ioutil.ReadFile("voice.data")
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	CallData := make([]VoiceData, 0)
	for _, v := range lines {
		call, ok := parseCall(v)
		if ok {
			CallData = append(CallData, call)
		}
	}
	fmt.Println("Calls: ", CallData)
}
func parseCall(line string) (VoiceData, bool) {
	call := strings.Split(line, ";")

	switch {
	case len(call) != 8:
		fallthrough
	case !isCountryOK(call[CountryCall]):
		fallthrough
	case !isProviderCallOK(call[ProviderCall]):
		return VoiceData{}, false
	}
	return VoiceData{
		Country:             call[CountryCall],
		Bandwidth:           call[BandwidthCall],
		ResponseTime:        call[ResponseTimeCall],
		Provider:            call[ProviderCall],
		ConnectionStability: call[ConnectionStabilityCall],
		TTFB:                call[TTFBCall],
		VoicePurity:         call[VoicePurityCall],
		MedianOfCallsTime:   call[MedianOfCallsTimeCall],
	}, true
}
