package util

import (
	"log"
	"os"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
