package conf

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"os"
)

// Описание конфига
type Config struct {
	Apport       string `yaml:"apport"`       // Порт
	Logfile      string `yaml:"logfile"`      // Файл логов
	Smsdata      string `yaml:"smsdata"`      // Файл данных смс
	Voicedata    string `yaml:"voicedata"`    // Файл данных звонков
	Incidentdata string `yaml:"incidentdata"` // Ссылка на API инцедентов
	Mmsdata      string `yaml:"mmsdata"`      // Ссылка на API ММS
	Supportdata  string `yaml:"supportdata"`  // Ссылка на API поддержки
	Emaildata    string `yaml:"emaildata"`    // Файл данных почты
	Billing      string `yaml:"billing"`      // Файл данных биллинга
}

var Cfg Config
var Con *Config

// Считываем настройки из файла

func GetConf(file string, cnf interface{}) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, cnf)
	}
	return err
}

// При запуске устанавливаем настройки логирования
func init() {

	if err := GetConf("my_conf.yml", &Cfg); err != nil {
		log.Panicln(err)
	}
	Con = &Cfg
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	// Если файл не указан в настройках то по умолчанию пишем в log.txt
	logfile := Con.Logfile
	if len([]rune(logfile)) == 0 {
		logfile = "log.txt"
	}
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(f)
	if err != nil {
		log.Fatal(err)
	}
}

// Функция для считывания IP из запроса
func ReadUserIP(req *http.Request) string {
	IPAddress := req.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = req.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = req.RemoteAddr
	}
	return IPAddress
}

// Функция для считывания UserAgent из запроса
func ReadUserAgent(req *http.Request) string {
	UserAgent := req.Header.Get("User-Agent")
	return UserAgent
}

// Записываем в лог файл запись о запросе
func LogRequestPayload(userIP, userAgent string) {
	log.Infof("userIP: %s, User-Agent: %s", userIP, userAgent)
}
