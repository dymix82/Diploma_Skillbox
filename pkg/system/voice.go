package system

import (
	"fmt"
	"io/ioutil"
	"strconv"
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
	ConnectionStability float32
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
	StabilityString := call[ConnectionStabilityCall]
	StabilityFloat64, _ := strconv.ParseFloat(StabilityString, 32)
	StabilityFloat32 := float32(StabilityFloat64)
	return VoiceData{
		Country:             call[CountryCall],
		Bandwidth:           call[BandwidthCall],
		ResponseTime:        call[ResponseTimeCall],
		Provider:            call[ProviderCall],
		ConnectionStability: StabilityFloat32,
		TTFB:                call[TTFBCall],
		VoicePurity:         call[VoicePurityCall],
		MedianOfCallsTime:   call[MedianOfCallsTimeCall],
	}, true
}
