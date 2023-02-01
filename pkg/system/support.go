package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"main/pkg/conf"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

// Загрузка данных о данных системы Поддержки
func ImportSupport() ([]SupportData, error) {
	response, err := http.Get(conf.Con.Supportdata)
	if err != nil {
		return []SupportData{}, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 || len(body) == 0 {
		err = errors.New("Не доступен API по загрузке информации по поддержке")
		fmt.Println(err)
		return []SupportData{}, err

	}
	response.Body.Close()
	SupportData := make([]SupportData, 0)
	json.Unmarshal(body, &SupportData)
	return SupportData, nil
}
