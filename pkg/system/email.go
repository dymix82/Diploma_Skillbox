package system

import (
	"io/ioutil"
	"main/pkg/conf"
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

// Читаем файл с данными о почте
func ImportEmail() ([]EmailData, error) {
	bytesRead, err := ioutil.ReadFile(conf.Con.Emaildata)
	if err != nil {
		return []EmailData{}, err
	}
	fileContent := string(bytesRead)
	lines := strings.Split(fileContent, "\n")
	Emaildata := make([]EmailData, 0)
	for _, v := range lines {
		mail, ok := parseMail(v)
		if ok {
			Emaildata = append(Emaildata, mail)
		}
	}
	// Сортируем по времени доставки
	sort.Slice(Emaildata, func(i, j int) bool {
		return Emaildata[i].DeliveryTime < Emaildata[j].DeliveryTime
	})
	return Emaildata, nil
}

// Функции проверки файла с данными о почте построчно
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
