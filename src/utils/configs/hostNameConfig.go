package configs

import (
	"omnenest-backend/src/constants"
	"os"
	"time"
)

var hostName string

// SetHostName sets the host name of the machine.
func SetHostName(service string) {
	var err error
	hostName, err = os.Hostname()
	if err != nil {
		ist, _ := time.LoadLocation(constants.IST)
		hostName = service + "-" + time.Now().In(ist).Format(constants.PrometheusDateTimeMinuteFormat)
	}
}

// GetHostName returns the host name of the machine.
func GetHostName() string {
	return hostName
}
