package log

import (
	"os"

	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("")
var debug bool
var formatConsole = logging.MustStringFormatter(
	`%{color}%{id:004d} %{time:15:04:05.000} %{module} %{shortfile} %{longfunc} ▶ %{level:.4s} %{color:reset} %{message}`,
)

var formatFile = logging.MustStringFormatter(
	`%{id:004d} %{time:15:04:05.000} %{module} %{shortfile} %{longfunc} ▶ %{level:.4s} %{message}`,
)

var console logging.Backend
var file logging.Backend
var filePath string

func init() {
	log.ExtraCalldepth = 1
	console = logging.NewLogBackend(os.Stderr, "", 0)
	console = logging.NewBackendFormatter(console, formatConsole)
	logging.SetBackend(console)
	if debug {
		logging.SetLevel(logging.DEBUG, "")
	} else {
		logging.SetLevel(logging.INFO, "")
	}
}

func DebugEanble(enable bool) {
	debug = enable
	if debug {
		logging.SetLevel(logging.DEBUG, "")
	} else {
		logging.SetLevel(logging.INFO, "")
	}
}

var f *os.File

func LogToFile(path string) {
	if filePath != path {
		f.Close()
	}
	filePath = path
	var err error
	f, err = os.OpenFile("foo.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Warning(err)
	}
	file = logging.NewLogBackend(f, "", 0)
	file = logging.NewBackendFormatter(file, formatFile)
	logging.SetBackend(console, file)
	if debug {
		logging.SetLevel(logging.DEBUG, "")
	} else {
		logging.SetLevel(logging.INFO, "")
	}
}

func getLevel() logging.Level {
	return logging.GetLevel("")
}

var Println = log.Debug
var Debug = log.Debug
var Info = log.Info
var Notice = log.Notice
var Warning = log.Warning
var Error = log.Error
var Fatal = log.Fatal
var Critical = log.Critical

var Printf = log.Debugf
var Debugf = log.Debugf
var Infof = log.Infof
var Noticef = log.Noticef
var Warningf = log.Warningf
var Errorf = log.Errorf
var Fatalf = log.Fatalf
var Criticalf = log.Criticalf
