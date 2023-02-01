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
	response, error := http.Get("http://127.0.0.1:8383/accendent")
	if error != nil {
		return []IncidentData{}, error
	}
	body, error := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 || len(body) == 0 {
		error = errors.New("Не доступен API по загрузке инцедентов")
		fmt.Println(error)
		return []IncidentData{}, error

	}
	response.Body.Close()
	IncidentData := make([]IncidentData, 0)
	json.Unmarshal(body, &IncidentData)
	return IncidentData, nil
}
