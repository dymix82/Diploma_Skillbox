package result

import "main/pkg/system"

type ResultT struct {
	Status bool       `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   ResultSetT `json:"data"`   // Заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}

type ResultSetT struct {
	SMS       [][]system.SMSData              `json:"sms"`
	MMS       [][]system.MMSData              `json:"mms"`
	VoiceCall []system.VoiceData              `json:"voice_call"`
	Email     map[string][][]system.EmailData `json:"email"`
	Billing   system.BillingData              `json:"billing"`
	Support   []int                           `json:"support"`
	Incidents []system.IncidentData           `json:"incident"`
}
