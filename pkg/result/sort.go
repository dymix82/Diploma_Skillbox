package result

import (
	"golang.org/x/exp/slices"
	"main/pkg/system"
	"sort"
)

func MakeSMSResult() ([][]system.SMSData, error) {
	smsSlicewithNames, err := system.ImportSMS()
	if err != nil {
		return [][]system.SMSData{}, err
	}
	for i := range smsSlicewithNames {
		//	if len([]rune(smsSlicewithNames[i].Country)) == 2 {
		smsSlicewithNames[i].Country = system.CodetoCountryname(smsSlicewithNames[i].Country)
		//	}
	}
	smsSortedbyCountry := make([]system.SMSData, len(smsSlicewithNames))
	copy(smsSortedbyCountry, smsSlicewithNames)
	sort.Slice(smsSortedbyCountry, func(i, j int) bool {
		return smsSortedbyCountry[i].Country < smsSortedbyCountry[j].Country
	})

	smsSortedbyProvider := make([]system.SMSData, len(smsSlicewithNames))
	copy(smsSortedbyProvider, smsSlicewithNames)
	sort.Slice(smsSortedbyProvider, func(i, j int) bool {
		return smsSortedbyProvider[i].Provider < smsSortedbyProvider[j].Provider
	})

	SMSResult := [][]system.SMSData{smsSortedbyProvider, smsSortedbyCountry}
	return SMSResult, nil
}
func MakeMMSResult() ([][]system.MMSData, error) {
	mmsSlicewithNames, error := system.ImportMMS()
	if error != nil {
		return [][]system.MMSData{}, error
	}
	for i := range mmsSlicewithNames {
		//	if len([]rune(mmsSlicewithNames[i].Country)) == 2 {
		mmsSlicewithNames[i].Country = system.CodetoCountryname(mmsSlicewithNames[i].Country)
		//	}
	}
	mmsSortedbyCountry := make([]system.MMSData, len(mmsSlicewithNames))
	copy(mmsSortedbyCountry, mmsSlicewithNames)
	sort.Slice(mmsSortedbyCountry, func(i, j int) bool {
		return mmsSortedbyCountry[i].Country < mmsSortedbyCountry[j].Country
	})

	mmsSortedbyProvider := make([]system.MMSData, len(mmsSlicewithNames))
	copy(mmsSortedbyProvider, mmsSlicewithNames)
	sort.Slice(mmsSortedbyProvider, func(i, j int) bool {
		return mmsSortedbyProvider[i].Provider < mmsSortedbyProvider[j].Provider
	})

	MMSResult := [][]system.MMSData{mmsSortedbyProvider, mmsSortedbyCountry}

	return MMSResult, nil

}

func MakeEmailResult() (map[string][][]system.EmailData, error) {
	MailResult := make(map[string][][]system.EmailData, 0)
	Emaildata, err := system.ImportEmail()
	if err != nil {
		return MailResult, err
	}
	CountrymapFast := make(map[string][]system.EmailData)
	countries := make([]string, 0)
	for _, v := range Emaildata {
		if !slices.Contains(countries, v.Country) {
			countries = append(countries, v.Country)
		}
	}

	count := 0
	for _, v := range countries {
		for i := range Emaildata {
			if v == Emaildata[i].Country {
				count++
				if count <= 3 {
					CountrymapFast[v] = append(CountrymapFast[v], Emaildata[i])
				} else {
					count = 0
					break
				}

			}

		}

	}
	CountrymapSlow := make(map[string][]system.EmailData)
	count = 0
	length := len(Emaildata) - 1
	for _, v := range countries {
		for i := range Emaildata {
			if v == Emaildata[length-i].Country {
				//fmt.Println(system.Emaildata[i])
				count++
				if count <= 3 {
					CountrymapSlow[v] = append(CountrymapSlow[v], Emaildata[length-i])
				} else {
					count = 0
					break
				}

			}

		}

	}

	for k := range CountrymapSlow {

		MailResult[k] = [][]system.EmailData{CountrymapFast[k], CountrymapSlow[k]}
	}
	return MailResult, nil
}
