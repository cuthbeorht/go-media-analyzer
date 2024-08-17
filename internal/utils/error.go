package utils

import "log"

func CheckError(e error) {
	log.Fatalf("%s", e)
}
