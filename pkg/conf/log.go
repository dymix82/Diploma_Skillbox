package conf

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
	"os"
)

type Config struct {
	Apport string `yaml:"apport"`
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

// При запуске устанавливаем настройки логгирования
func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)
	if err := GetConf("my_conf.yml", &Cfg); err != nil {
		log.Panicln(err)
	}
	Con = &Cfg
	Con = &Cfg
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
