package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

func ImportSupport() {
	response, error := http.Get("http://localhost:8383/support")
	if error != nil {
		log.Fatalln()
	}
	body, error := ioutil.ReadAll(response.Body)
	response.Body.Close()
	SupportData := make([]SupportData, 0)
	json.Unmarshal(body, &SupportData)
	fmt.Println(SupportData)
}
