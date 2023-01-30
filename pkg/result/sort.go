package result

import (
	"main/pkg/system"
	"sort"
)

func MakeSMSResult() [][]system.SMSData {
	smsSlicewithNames := system.SMSdata
	for i := range smsSlicewithNames {
		//	if len([]rune(smsSlicewithNames[i].Country)) == 2 {
		smsSlicewithNames[i].Country = system.CodetoCountryname(smsSlicewithNames[i].Country)
		//	}
	}
	smsSortedbyCountry := make([]system.SMSData, len(system.SMSdata))
	copy(smsSortedbyCountry, smsSlicewithNames)
	sort.Slice(smsSortedbyCountry, func(i, j int) bool {
		return smsSortedbyCountry[i].Country < smsSortedbyCountry[j].Country
	})

	smsSortedbyProvider := make([]system.SMSData, len(system.SMSdata))
	copy(smsSortedbyProvider, smsSlicewithNames)
	sort.Slice(smsSortedbyProvider, func(i, j int) bool {
		return smsSortedbyProvider[i].Provider < smsSortedbyProvider[j].Provider
	})

	SMSResult := [][]system.SMSData{smsSortedbyCountry, smsSortedbyProvider}
	return SMSResult
}
func MakeMMSResult() [][]system.MMSData {
	mmsSlicewithNames := system.MMSdataFiltred
	for i := range mmsSlicewithNames {
		//	if len([]rune(mmsSlicewithNames[i].Country)) == 2 {
		mmsSlicewithNames[i].Country = system.CodetoCountryname(mmsSlicewithNames[i].Country)
		//	}
	}
	mmsSortedbyCountry := make([]system.MMSData, len(system.SMSdata))
	copy(mmsSortedbyCountry, mmsSlicewithNames)
	sort.Slice(mmsSortedbyCountry, func(i, j int) bool {
		return mmsSortedbyCountry[i].Country < mmsSortedbyCountry[j].Country
	})

	mmsSortedbyProvider := make([]system.MMSData, len(system.SMSdata))
	copy(mmsSortedbyProvider, mmsSlicewithNames)
	sort.Slice(mmsSortedbyProvider, func(i, j int) bool {
		return mmsSortedbyProvider[i].Provider < mmsSortedbyProvider[j].Provider
	})

	MMSResult := [][]system.MMSData{mmsSortedbyCountry, mmsSortedbyProvider}

	return MMSResult

}
