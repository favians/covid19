package cron

import (
	"log"

	"github.com/favians/golang_starter/bootstrap"
)

func SendSMS() {
	log.Println("Send SMS to Client")
	bootstrap.App.Log.Logger.Println("Done to sending SMS")
}
