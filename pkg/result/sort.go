package result

import (
	"golang.org/x/exp/slices"
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

	SMSResult := [][]system.SMSData{smsSortedbyProvider, smsSortedbyCountry}
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

	MMSResult := [][]system.MMSData{mmsSortedbyProvider, mmsSortedbyCountry}

	return MMSResult

}

func MakeEmailResult() map[string][][]system.EmailData {
	coutries := make([]string, 0)
	for _, v := range system.Emaildata {
		if !slices.Contains(coutries, v.Country) {
			coutries = append(coutries, v.Country)
		}
	}

	CountrymapFast := make(map[string][]system.EmailData)
	count := 0
	for _, v := range coutries {
		for i := range system.Emaildata {
			if v == system.Emaildata[i].Country {
				count++
				if count <= 3 {
					CountrymapFast[v] = append(CountrymapFast[v], system.Emaildata[i])
				} else {
					count = 0
					break
				}

			}

		}

	}
	CountrymapSlow := make(map[string][]system.EmailData)
	count = 0
	length := len(system.Emaildata) - 1
	for _, v := range coutries {
		for i := range system.Emaildata {
			if v == system.Emaildata[length-i].Country {
				//fmt.Println(system.Emaildata[i])
				count++
				if count <= 3 {
					CountrymapSlow[v] = append(CountrymapSlow[v], system.Emaildata[length-i])
				} else {
					count = 0
					break
				}

			}

		}

	}

	MailResult := make(map[string][][]system.EmailData, 0)
	for k := range CountrymapSlow {

		MailResult[k] = [][]system.EmailData{CountrymapFast[k], CountrymapSlow[k]}
	}
	return MailResult
}
