package utils

import "log"

func ErrorPanics(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorLogs(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
