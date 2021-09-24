package log

import (
	"net/http"
	"os"
	"time"

	"github.com/addonrizky/sagaracrud/config"
	"github.com/addonrizky/sagaracrud/constant"
	"github.com/addonrizky/sagaracrud/entity/entityself"

	log "github.com/sirupsen/logrus"
)

var (
	serviceName string
)

const (
	httpRequest  = "REQUEST"
	httpResponse = "RESPONSE"
	timeformat   = "2006-01-02T15:04:05-0700"
	nameformat   = "log-2006-01-02.log"
)

func LogDebug(msg string) {
	setText()
	timestamp := setLogFile()
	log.Debug(timestamp + " " + msg)
}
func LogWarn(err, location string) {
	setText()
	timestamp := setLogFile()
	log.WithFields(log.Fields{
		"service":        serviceName,
		"error":          err,
		"error_location": location,
		"timestamp":      timestamp,
	}).Warn("WARNING")
}

func LogResponse(username, trx_type, code string, response *entityself.Response, header http.Header) {
	setJSON()
	timestamp := setLogFile()
	log.WithFields(log.Fields{
		"service":         serviceName,
		"http_type":       httpResponse,
		"response_header": header,
		"response_body":   response,
		"trx_type":        trx_type,
		"username":        username,
		"response_code":   code,
		"timestamp":       timestamp,
	}).Info("RESPONSE")
}

func LogRequest(username, trx_type string, request interface{}, header http.Header) {
	setJSON()
	timestamp := setLogFile()
	log.WithFields(log.Fields{
		"service":        serviceName,
		"http_type":      httpRequest,
		"request_header": header,
		"request_body":   request,
		"trx_type":       trx_type,
		"username":       username,
		"timestamp":      timestamp,
	}).Info("REQUEST")

}

func setLogFile() string {
	currentTime := time.Now()
	timestamp := currentTime.Format(timeformat)
	filename := "logs/" + currentTime.Format(nameformat)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		LogWarn(err.Error(), "set logfile")
	} else {
		log.SetOutput(file)
	}
	return timestamp
}

func setJSON() {
	formatter := new(log.JSONFormatter)
	formatter.DisableTimestamp = true
	log.SetFormatter(formatter)
}

func setText() {
	formatter := new(log.TextFormatter)
	formatter.DisableTimestamp = true
	formatter.DisableQuote = true
	log.SetFormatter(formatter)
}

func Init() {
	if config.GetBool("DEBUG") {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	formatter := new(log.JSONFormatter)
	formatter.DisableTimestamp = true
	log.SetFormatter(formatter)

	serviceName = constant.ServiceName
}
