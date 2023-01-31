package system

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

const (
	CountryEmail = iota
	ProviderEmail
	DeliveryTime
)

var Emaildata []EmailData

func ImportEmail() {
	bytesRead, _ := ioutil.ReadFile("email.data")
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	Emaildata = make([]EmailData, 0)
	for _, v := range lines {
		mail, ok := parseMail(v)
		if ok {
			Emaildata = append(Emaildata, mail)
		}
	}
	sort.Slice(Emaildata, func(i, j int) bool {
		return Emaildata[i].DeliveryTime < Emaildata[j].DeliveryTime
	})
}
func parseMail(line string) (EmailData, bool) {
	mail := strings.Split(line, ";")

	switch {
	case len(mail) != 3:
		fallthrough
	case !isCountryOK(mail[CountryEmail]):
		fallthrough
	case !isProviderMailOK(mail[ProviderEmail]):
		return EmailData{}, false
	}
	DeliveryTimeInt, err := strconv.Atoi(mail[DeliveryTime])
	if err != nil {
		return EmailData{}, false
	}
	return EmailData{
		Country:      mail[CountryEmail],
		Provider:     mail[ProviderEmail],
		DeliveryTime: DeliveryTimeInt,
	}, true
}
