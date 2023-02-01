package system

import (
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
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                string  `json:"ttfb"`
	VoicePurity         string  `json:"voice_purity"`
	MedianOfCallsTime   string  `json:"median_of_calls_time"`
}

func ImportVoice() ([]VoiceData, error) {
	bytesRead, err := ioutil.ReadFile("voice.data")
	if err != nil {
		return []VoiceData{}, err
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	callData := make([]VoiceData, 0)
	for _, v := range lines {
		call, ok := parseCall(v)
		if ok {
			callData = append(callData, call)
		}
	}
	return callData, nil
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
