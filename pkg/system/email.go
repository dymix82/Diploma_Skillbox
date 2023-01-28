package system

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}

const (
	CountryEmail = iota
	ProviderEmail
	DeliveryTime
)

func ImportEmail() {
	bytesRead, _ := ioutil.ReadFile("email.data")
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	Emaildata := make([]EmailData, 0)
	for _, v := range lines {
		mail, ok := parseMail(v)
		if ok {
			Emaildata = append(Emaildata, mail)
		}
	}
	fmt.Println("Email: ", Emaildata)
	//for i, _ := range SMSdata {
	//	fmt.Println(SMSdata[i])
	//}

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
