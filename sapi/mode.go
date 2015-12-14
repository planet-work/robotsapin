package sapi

import (
	"io/ioutil"
	"os"
)

const EnvSapiLogMode = "SAPI_LOG_MODE"

const (
	DebugLogMode string = "debug"
	InfoLogMode  string = "info"
	ErrorLogMode string = "error"
	TestDbMode   string = "test"
	ProdDbMode   string = "prod"
)

var logModeName = DebugLogMode
var dbModeName = ProdDbMode

func DetectLogMode() {
	logMode := os.Getenv(EnvSapiLogMode)
	if len(logMode) == 0 {
		SetLogMode(DebugLogMode)
	} else {
		SetLogMode(logMode)
	}
}

func SetLogMode(value string) {
	switch value {
	case DebugLogMode:
		initLogger(os.Stdout, os.Stdout, os.Stderr)
	case InfoLogMode:
		initLogger(ioutil.Discard, os.Stdout, os.Stderr)
	case ErrorLogMode:
		initLogger(ioutil.Discard, ioutil.Discard, os.Stderr)
	default:
		panic("sapi log mode unknown: " + value)
	}
	logModeName = value
}

func SetDbMode(value string) {
	switch value {
	case TestDbMode:
		dbModeName = value
	case ProdDbMode:
		dbModeName = value
	default:
		panic("sapi db mode unknown: " + value)
	}
}

func LogMode() string {
	return logModeName
}

func DbMode() string {
	return dbModeName
}
