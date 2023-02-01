package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"main/pkg/conf"
	"net/http"
)

var MMSdataFiltred []MMSData

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

// Импорт  данных о ММС через API
func ImportMMS() ([]MMSData, error) {
	// make GET request
	response, err := http.Get(conf.Con.Mmsdata)
	if err != nil {
		return []MMSData{}, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 || len(body) == 0 {
		err = errors.New("Не доступен API по загрузке MMS")
		fmt.Println(err)
		return []MMSData{}, err

	}
	response.Body.Close()
	MMSdata := make([]MMSData, 0)
	json.Unmarshal(body, &MMSdata)
	MMSdataFiltred = make([]MMSData, 0)
	for i := range MMSdata {
		switch {
		case !isCountryOK(MMSdata[i].Country):
			fallthrough
		case !isProviderSMSOK(MMSdata[i].Provider):
			continue
		}
		MMSdataFiltred = append(MMSdataFiltred, MMSdata[i])
	}
	return MMSdataFiltred, nil
}
