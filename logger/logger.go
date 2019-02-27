package logger

import (
	"github.com/op/go-logging"
	"os"
)

var MyLogger = logging.MustGetLogger("example")

// Example format string. Everything except the message has a custom color
// which is dependent on the MyLogger level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:02x}%{color:reset} %{message}`,
)
var format3 = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:02x}%{color:reset} %{message}`,
)

// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

//
//func main() {
//
//	MyLogger.Debugf("debug %s  _____________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________", Password("secret"))
//	MyLogger.Info("info _____________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________")
//	MyLogger.Notice("notice _____________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________")
//	MyLogger.Warning("warning _____________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________")
//	MyLogger.Error("err _____________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________")
//	MyLogger.Critical("crit _____________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________________")
//}

func init() {
	// For demo purposes, create two backend for os.Stderr.
	//backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	//backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	backend3 := logging.NewLogBackend(os.Stdout, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used MyLogger level and the name of
	// the function.
	//backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend3Formatter := logging.NewBackendFormatter(backend3, format3)

	// Only errors and more severe messages should be sent to backend1
	//backend1Leveled := logging.AddModuleLevel(backend1)
	//backend1Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	//logging.SetBackend(backend1Leveled, backend2Formatter)
	logging.SetBackend(backend3Formatter)
	//logging.SetBackend(backend2Formatter)
}
