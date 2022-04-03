package logger

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"inshortsProj/constant"
	"os"
)

var InfoLog = log.New()
var ErrorLog = log.New()
var DebugLog = log.New()

func init() {
	///////////////////////Check for log file and create if not there//////////////////////////////
	_, err := os.OpenFile(constant.LogInfoFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic("Failed to create info log file: " + err.Error())
	}
	_, err = os.OpenFile(constant.LogAccessFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic("Failed to create access log file: " + err.Error())
	}
	FilePath := constant.LogInfoFilePath
	InfoObject := &lumberjack.Logger{
		Filename:   FilePath,
		MaxSize:    2, // megabytes
		MaxBackups: 3,
		Compress:   true,
	}
	//////////////////////////InfoLogs Settings/////////////////////////////
	InfoLog.Formatter = new(log.JSONFormatter)
	InfoLog.Level = log.InfoLevel // setting this to lowest so that all logs will be pushed.
	InfoLog.Out = InfoObject
	//InfoLog.Out = os.Stdout //// To be Removed

	//////////////////////////ErrorLogs Settings/////////////////////////////
	ErrorLog.Formatter = new(log.JSONFormatter)
	ErrorLog.Level = log.ErrorLevel
	ErrorLog.Out = InfoObject
	//ErrorLog.Out = os.Stdout // tp be removed

	DebugLog.Formatter = new(log.JSONFormatter)
	DebugLog.Level = log.DebugLevel
	DebugLog.Out = os.Stdout
}

func LogInfoForScalyr(Description string, FunctionNameDescription string, vertical string, unique_id_tagging interface{}) {
	MobileCore := constant.MOBILECOREAPP
	InfoLog.WithFields(log.Fields{
		"Log":         MobileCore,
		"Description": Description,
		"Block":       FunctionNameDescription,
		"Vertical":    vertical,
		"tag":         unique_id_tagging,
	}).Info("Info logs")
}

func LogErrorForScalyr(Error string, FunctionNameDescription string, vertical string, unique_id_tagging interface{}) {
	MobileCore := constant.MOBILECOREAPP
	ErrorLog.WithFields(log.Fields{
		"Log":      MobileCore,
		"Error":    Error,
		"Block":    FunctionNameDescription,
		"Vertical": vertical,
		"tag":      unique_id_tagging,
	}).Error("Error logs")
}

func Debug(Description string, FunctionNameDescription string, vertical string, unique_id_tagging interface{}) {

	MobileCore := constant.MOBILECOREAPP
	DebugLog.WithFields(log.Fields{
		"Log":         MobileCore,
		"Description": Description,
		"Block":       FunctionNameDescription,
		"Vertical":    vertical,
		"tag":         unique_id_tagging,
	}).Debug("Debug logs \n ")
}
