package logging

import (
	"github.com/sirupsen/logrus"
	"os"
	"fmt"
	"errors"
)

// main initializes default logs ve calling func CreateDefaultLogs (). three logs will be created
// DEFAULT : for bugs, panics, errors, warnings
// DEBUG  : for debuging purposes ie variable values, flow checks
// INFO : Something important happened
// PERF : Timing, durations etc.
// Modules can get the default loggers mylog := logging.LogMan.GetLogger ("DEFAULT"). 
// They can also create their own log as
// logman := new (logging.LogMan)
// mylog, err := logman.NewLogger (logname, logfilepath)

type LogMan struct {
	logs map[string]*logrus.Logger
}
var logman LogMan


func NewLogger (logname string, logpath string)  (*logrus.Logger, error) {
	var newlog *logrus.Logger

	if logman.logs == nil {
		logman.logs = make (map[string]*logrus.Logger)
	}

	switch {
	case logname == ""  && logpath == "" :

		if existinglog, ok := logman.logs["DEFAULT"]; ok { // There is already a DEFAULT, then return it
			return existinglog, nil
		}

		newlog = logrus.New ()
		if newlog == nil {
			return nil, fmt.Errorf ("NewLogger: Can not crete logrus log %v", logname)
		} else {
			newlog.SetFormatter(&logrus.JSONFormatter{})
			logman.logs["DEFAULT"]  = newlog
			return logman.logs["DEFAULT"], nil
		}
	case logname == "" && logpath != "" :
		if existinglog, ok := logman.logs["DEFAULT"]; ok { // There is already a DEFAULT, then return it
			return existinglog, nil
		}

		newlog = logrus.New ()
		if newlog == nil {
			return nil, fmt.Errorf ("NewLogger: Can not crete logrus log with path %v", logpath)
		} else {
			file, err := os.OpenFile(logpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err == nil {
				newlog.SetFormatter(&logrus.JSONFormatter{})
				newlog.Out = file
				logman.logs["DEFAULT"] = newlog
				return logman.logs["DEFAULT"], nil
			} else {
				return nil, fmt.Errorf ("NewLogger: Can not crete log file %v", logpath)
			}
		}
	case logname != "" && logpath != "" :
		if existinglog, ok := logman.logs[logname]; ok { // There is already a log in that name, then return it
			return existinglog, nil
		}

		newlog = logrus.New ()
		if newlog == nil {
			return nil, fmt.Errorf ("NewLogger: Can not crete logrus log with path %v and name %v", logpath, logname)
		} else {
			file, err := os.OpenFile(logpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err == nil {
				newlog.SetFormatter(&logrus.JSONFormatter{})
				newlog.Out = file
				logman.logs[logname] = newlog
				return logman.logs[logname], nil
			} else {
				return nil, fmt.Errorf ("NewLogger: Can not crete log file %v and name %v", logpath, logname)
			}
		}
	case logname != "" && logpath == "" :
		if existinglog, ok := logman.logs[logname]; ok { // There is already a log in that name, then return it
			fmt.Println ("Exitsting log var")
			return existinglog, nil
		}

		newlog = logrus.New ()
		if newlog == nil {
			return nil, fmt.Errorf ("NewLogger: Can not crete logrus log without path %v", logname)
		} else {
			newlog.SetFormatter(&logrus.JSONFormatter{})
			logman.logs[logname] = newlog
			return logman.logs[logname], nil
		}
	default :
		return nil, nil
	}


	return nil, nil // should not be
}

func GetLogger (logname string)  (*logrus.Logger, error) {
	if logname == "" {
		if logman.logs["DEFAULT"] == nil {
			return nil, fmt.Errorf("GetLogger : No default log entry found ")
		} else {
			return logman.logs ["DEFAULT"], nil
		}
	} else {
		if logman.logs[logname] == nil {
			return nil, fmt.Errorf("GetLogger : No log entry found %v ", logname)
		} else {
			return logman.logs [logname], nil
		}
	}
}

// DEFAULT : for bugs, panics, errors, warnings
// DEBUG  : for debuging purposes ie variable values, flow checks
// INFO : Something important happened
// PERF : Timing, durations etc.
func CreateDefaultLogs () error {
	if tmplog, err := NewLogger ("DEFAULT", ""); err != nil {
                fmt.Println ("DEFAULT log can not be initialized")
                return errors.New ("CreateDefaultLogs ; Can not create DEFAULT log")
	} else {
		logman.logs["DEFAULT"] = tmplog
		fmt.Println ("DEFAULT yaratildi")
        }
	if tmplog, err := NewLogger ("DEBUG", "./logs/debug.log"); err != nil {
                fmt.Println ("DEBUG log can not be initialized")
                return errors.New ("CreateDefaultLogs ; Can not create DEBUG log")
	} else {
		logman.logs["DEBUG"] = tmplog
		fmt.Println ("DEBUG yaratildi")
        }
	if tmplog, err := NewLogger ("INFO", "./logs/info.log"); err != nil {
                fmt.Println ("INFO log can not be initialized")
                return errors.New ("CreateDefaultLogs ; Can not create INFO log")
	} else {
		logman.logs["INFO"] = tmplog
		fmt.Println ("INFO yaratildi")
        }
	if tmplog, err := NewLogger ("PERF", "./logs/perf.log"); err != nil {
                fmt.Println ("PERF log can not be initialized")
                return errors.New ("CreateDefaultLogs ; Can not create PERF log")
	} else {
		logman.logs["PERF"] = tmplog
		fmt.Println ("PERF yaratildi")
        }

	fmt.Println ("Logs : ", logman.logs)
	return nil
}
