package cron

import (
	"log"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/favians/golang_starter/bootstrap"
)

type (
	RSresult struct {
		ID           int       `json:"id"`
		Nama         string    `json:"nama"`
		Lower        int       `json:"lower"`
		Upper        int       `json:"upper"`
		Start        time.Time `json:"start"`
		Stop         time.Time `json:"stop"`
		NextSchedule time.Time `json:"next_schedule"`
	}

	PasienResult struct {
		NoHP string `json:"no_hp"`
		Kode string `json:"kode"`
	}
)

type CronJob struct {
	// struct difiner
}

func (e CronJob) Run() {
	DoCron()
}

func DoCron() {
	GetRumahSakit()
	up := bootstrap.App.AppConfig.GetSub("sms_interval").GetInt("upper")
	low := bootstrap.App.AppConfig.GetSub("sms_interval").GetInt("lower")
	start := bootstrap.App.AppConfig.GetSub("sms_interval").GetString("start")
	stop := bootstrap.App.AppConfig.GetSub("sms_interval").GetString("stop")
	random := rand.Intn(up-low) + low

	bootstrap.App.Log.Logger.Println("Cron Job Started in" + string(random))

	//Register Your Logic Function Here
	SendTime := isInTimeRange(start, stop)

	if SendTime {
		bootstrap.App.Log.Logger.Println("Trying to sending SMS")
		SendSMS()
	}
}

func GetRumahSakit() {
	qres := []RSresult{}
	qry := bootstrap.App.DB.Table("rumah_sakits").Where("next_schedule < NOW()").Select("id, nama, lower, upper, start, stop, next_schedule")
	qry.Scan(&qres)

	for _, value := range qres {
		GetPasien(value.ID)
		updateSchedule(qry, value)
	}
}

func GetPasien(rsID int) {
	pasien := []PasienResult{}

	bootstrap.App.DB.Table("pasiens").Where("pasiens.rumah_sakit_id = ?", rsID).Select("no_hp, kode").Scan(&pasien)
	for _, value := range pasien {
		log.Println(value.NoHP)
		log.Println(value.Kode)
	}
}

func updateSchedule(qry *gorm.DB, qres RSresult) {
	// random := rand.Intn(qres.Upper-qres.Lower) + qres.Lower

	// qry.Update()
}

func isInTimeRange(started string, stopped string) bool {

	t := time.Now()
	timeNowString := t.Format(time.Kitchen)
	timeNow := stringToTime(timeNowString)
	start := stringToTime(started)
	end := stringToTime(stopped)

	if timeNow.Before(start) {
		return false
	}

	if timeNow.Before(end) {
		return true
	}

	return false
}

func stringToTime(str string) time.Time {
	tm, err := time.Parse(time.Kitchen, str)
	if err != nil {
		bootstrap.App.Log.Logger.Println("Failed to decode time:", err.Error())
	}
	bootstrap.App.Log.Logger.Println("Time decoded:", tm)
	return tm
}
