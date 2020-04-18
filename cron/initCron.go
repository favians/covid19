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
		Start        string    `json:"start"`
		Stop         string    `json:"stop"`
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
	up := bootstrap.App.AppConfig.GetSub("sms_interval").GetInt("upper")
	low := bootstrap.App.AppConfig.GetSub("sms_interval").GetInt("lower")
	random := rand.Intn(up-low) + low

	bootstrap.App.Log.Logger.Println("Cron Job Started in" + string(random))
	GetRumahSakit()
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
		SendSMS(value)
	}
}

func updateSchedule(qry *gorm.DB, qres RSresult) {
	times := time.Now()
	random := rand.Intn(qres.Upper-qres.Lower) + qres.Lower
	nextSchedule := qres.NextSchedule.Add(time.Hour * time.Duration(random))

	if isInTimeRange(qres.Start, qres.Stop, nextSchedule) {
		qry.Where("id = ?", qres.ID).Update("next_schedule", nextSchedule)
	} else {
		t, err := time.Parse("03:04PM", qres.Stop)
		if err != nil {
			bootstrap.App.Log.Logger.Println("cron:initCron:updateSchedule() error in parsing Time")
		}
		selisih := Abs(t.Hour() - times.Hour())
		nextSchedule := times.Add(time.Hour * time.Duration(selisih))
		log.Println(nextSchedule)
		qry.Where("id = ?", qres.ID).Update("next_schedule", nextSchedule)
	}
}

func isInTimeRange(started string, stopped string, nextSchedule time.Time) bool {

	t := nextSchedule
	NextSchedulerString := t.Format(time.Kitchen)
	nextScheduler := stringToTime(NextSchedulerString)
	start := stringToTime(started)
	end := stringToTime(stopped)

	if nextScheduler.Before(start) {
		return false
	}

	if nextScheduler.Before(end) {
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
