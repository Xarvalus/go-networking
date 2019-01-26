package utils

import "log"

// TODO: better way to handle errors?

func LogFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
