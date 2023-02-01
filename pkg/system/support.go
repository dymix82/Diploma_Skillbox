package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func ImportSupport() ([]SupportData, error) {
	response, error := http.Get("http://localhost:8383/support")
	if error != nil {
		return []SupportData{}, error
	}
	body, error := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 || len(body) == 0 {
		error = errors.New("Не доступен API по загрузке информации по поддержке")
		fmt.Println(error)
		return []SupportData{}, error

	}
	response.Body.Close()
	SupportData := make([]SupportData, 0)
	json.Unmarshal(body, &SupportData)
	fmt.Println(SupportData)
	return SupportData, nil
}
