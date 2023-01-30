package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

func ImportIncident() {
	response, error := http.Get("http://127.0.0.1:8383/accendent")
	if error != nil {
		log.Fatalln()
	}
	body, error := ioutil.ReadAll(response.Body)
	response.Body.Close()
	IncidentData := make([]IncidentData, 0)
	json.Unmarshal(body, &IncidentData)
	fmt.Println(IncidentData)
}
