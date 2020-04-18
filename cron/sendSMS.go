package cron

import (
	"log"

	"github.com/favians/golang_starter/bootstrap"
)

func SendSMS(pasien PasienResult) {

	bootstrap.App.Log.Logger.Println("Sending SMS to:" + pasien.NoHP)
	bootstrap.App.Log.Logger.Println("Code:" + pasien.Kode)
	log.Println(pasien.NoHP)
	log.Println(pasien.Kode)
	bootstrap.App.Log.Logger.Println("Done to sending SMS")
}
