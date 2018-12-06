package log

import (
	"os"

	"net/http"

	"github.com/cryptix/exp/wslog"
	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("")
var debug = true
var formatConsole = logging.MustStringFormatter(
	`%{color}%{id:004d} %{time:2006-01-02 15:04:05.0000000} %{module} %{shortfile} %{longfunc} ▶ %{level:.4s} %{color:reset} %{message}`,
)

var formatFile = logging.MustStringFormatter(
	`%{id:004d} %{time:2006-01-02 15:04:05.0000000} %{module} %{shortfile} %{longfunc} ▶ %{level:.4s} %{message}`,
)

var ws logging.Backend

type Backends struct {
	activeBackends []logging.Backend
}

func (b *Backends) Add(e logging.Backend) {
	b.activeBackends = append(b.activeBackends, e)
	logging.SetBackend(b.activeBackends...)
	if debug {
		logging.SetLevel(logging.DEBUG, "")
	} else {
		logging.SetLevel(logging.INFO, "")
	}
}

var bks Backends

func init() {
	var console logging.Backend
	//add 1 call depth
	log.ExtraCalldepth = 1
	console = logging.NewLogBackend(os.Stderr, "", 0)
	console = logging.NewBackendFormatter(console, formatConsole)
	bks.Add(console)
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

func LogToFile(path string) {
	var file logging.Backend
	var f *os.File
	var err error
	f, err = os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Warning(err)
	}
	file = logging.NewLogBackend(f, "", 0)
	file = logging.NewBackendFormatter(file, formatFile)
	bks.Add(file)
	if debug {
		logging.SetLevel(logging.DEBUG, "")
	} else {
		logging.SetLevel(logging.INFO, "")
	}
}

func LogToWs(addr, path string) {
	var websocket *wslog.WebsocketBackend
	websocket = wslog.NewBackend()
	//ws := logging.NewBackendFormatter(websocket, formatFile)
	bks.Add(websocket)
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			//allow cros
			r.Header["Origin"] = []string{}
			websocket.ServeHTTP(w, r)

		})
		http.ListenAndServe(addr, mux)
	}()
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
var Panic = log.Panic

var Printf = log.Debugf
var Debugf = log.Debugf
var Infof = log.Infof
var Noticef = log.Noticef
var Warningf = log.Warningf
var Errorf = log.Errorf
var Fatalf = log.Fatalf
var Criticalf = log.Criticalf
var Panicf = log.Panicf
