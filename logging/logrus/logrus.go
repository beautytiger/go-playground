package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func mainFinal() {
	fmt.Println("main exits, finally")
}

func initLogrus() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true)
	fileWriter, err := os.OpenFile("logrus.log", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("error open log file")
	}
	logrus.SetOutput(io.MultiWriter(os.Stdout, fileWriter))
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	defer mainFinal()
	initLogrus()
	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	//logrus.Panic("panic msg")
	//logrus.Fatal("fatal msg")

	logrus.WithFields(logrus.Fields{
		"name": "roby",
		"age": 18,
		"fake": true,
	}).Info("with field")

	fieldLogger := logrus.WithFields(logrus.Fields{
		"name": "fieldLogger",
		"age": 18,
		"fake": true,
	})

	fieldLogger.Info("hello world.")
}