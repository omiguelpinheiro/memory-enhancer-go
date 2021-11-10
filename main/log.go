package main

import (
	"log"

	"github.com/go-errors/errors"
)

func InitLog() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func check(err error) {
	if err != nil {
		errors.Errorf(err.Error())
	}
}