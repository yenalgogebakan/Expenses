package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"logging"
)
// Global variables
var mainlog, infolog *logrus.Logger

func main () {
	// Log setup
	if err := logging.CreateDefaultLogs (); err != nil {
		fmt.Println ("Can not initialize logs, exiting")
		return
	}
	if tmplog, err := logging.GetLogger ("DEFAULT"); err != nil {
		fmt.Println ("main DEFAULT log can not be accessed, exiting")
		return
	} else {
		mainlog = tmplog
	}
	if tmplog, err := logging.GetLogger ("INFO"); err != nil {
		fmt.Println ("main INFO log can not be accessed, exiting")
		return
	} else {
		infolog = tmplog
	}
	infolog.WithFields(logrus.Fields{
		"package" : "main",
		"source": "main.go",
		"func": "main ()",
	}).Info("Expenses is starting !!!")

}
