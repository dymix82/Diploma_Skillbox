package result

import (
	log "github.com/sirupsen/logrus"
	"main/pkg/system"
)

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

// Создание структуры c результатами для вывода
func MakeStruct() ResultT {
	status := true
	var ShowedErr string

	var err error
	var resultSetT ResultSetT
	var resultT ResultT

	resultSetT.MMS, err = MakeMMSResult()
	if err != nil {
		log.Error(err)
		status = false
		ShowedErr = err.Error() + "\t"
	}

	resultSetT.SMS, err = MakeSMSResult()
	if err != nil {
		log.Error(err)
		status = false
		ShowedErr += err.Error() + "\t"
	}
	resultSetT.VoiceCall, err = system.ImportVoice()
	if err != nil {
		log.Error(err)
		status = false
		ShowedErr += err.Error() + "\t"
	}
	resultSetT.Billing, err = system.ImportBilling()
	if err != nil {
		log.Error(err)
		status = false
		ShowedErr += err.Error() + "\t"
	}
	resultSetT.Email, err = MakeEmailResult()
	if err != nil {
		log.Error(err)
		status = false
		ShowedErr += err.Error() + "\t"
	}
	resultSetT.Incidents, err = system.ImportIncident()
	if err != nil {
		log.Error(err)
		status = false
		ShowedErr += err.Error() + "\t"
	}

	resultSetT.Support, err = getSupport()
	if err != nil {
		log.Error(err)
		status = false
		ShowedErr += err.Error() + "\t"
	}

	resultT.Status = status
	resultT.Data = resultSetT
	resultT.Error = ShowedErr

	return resultT
}
