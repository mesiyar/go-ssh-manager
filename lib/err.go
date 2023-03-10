package lib

import "log"

func HandlerErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s error: %v", msg, err)

	}
}
