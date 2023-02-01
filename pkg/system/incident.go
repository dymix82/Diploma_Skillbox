package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

func ImportIncident() ([]IncidentData, error) {
	response, err := http.Get("http://127.0.0.1:8383/accendent")
	if err != nil {
		return []IncidentData{}, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 || len(body) == 0 {
		err = errors.New("Не доступен API по загрузке инцедентов")
		fmt.Println(err)
		return []IncidentData{}, err

	}
	response.Body.Close()
	IncidentData := make([]IncidentData, 0)
	json.Unmarshal(body, &IncidentData)
	return IncidentData, nil
}
