package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var MMSdataFiltred []MMSData

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

func ImportMMS() ([]MMSData, error) {
	// make GET request
	response, error := http.Get("http://localhost:8383/mms")
	if error != nil {
		return []MMSData{}, error
	}

	body, error := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 || len(body) == 0 {
		error = errors.New("Не доступен API по загрузке MMS")
		fmt.Println(error)
		return []MMSData{}, error

	}
	response.Body.Close()
	MMSdata := make([]MMSData, 0)
	json.Unmarshal(body, &MMSdata)
	MMSdataFiltred = make([]MMSData, 0)
	for i, _ := range MMSdata {
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
