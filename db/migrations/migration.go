package migrations

import (
	"encoding/json"
	"log"

	"github.com/favians/golang_starter/api/models"
	"github.com/favians/golang_starter/bootstrap"
)

// Attention on table relation, seed table that not mutual first.
func Seed() {
	if bootstrap.App.ENV == "dev" || bootstrap.App.ENV == "test" {

		seeder := bootstrap.App.DB.Begin()

		log.Printf("------------- SEED BEGIN ---------------")

		log.Printf("------------- SEED User ---------------")

		var admins []models.Admin = []models.Admin{
			models.Admin{Name: "vian", Username: "vian", Password: "1234", Instansi: "Rs. Mardi Waluyo"},
			models.Admin{Name: "arafat", Username: "arafat", Password: "1234", Instansi: "Rs. Syaiful Anwar"},
			models.Admin{Name: "hasan", Username: "hasan", Password: "1234", Instansi: "Rs. Persahabatan"},
			models.Admin{Name: "fikri", Username: "fikri", Password: "1234", Instansi: "Rs. Perjalanan Tetangga"},
		}

		for _, admin := range admins {
			if err := seeder.Create(&admin).Error; err != nil {
				debugLog, _ := json.Marshal(err)
				log.Printf(string(debugLog))
			} else {
				log.Printf("created")
			}
		}

		var RumahSakits []models.RumahSakit = []models.RumahSakit{
			models.RumahSakit{Nama: "Rs. Persahabatan", Start: "01:00AM", Stop: "11:00PM", Lower: "1", Upper: "5"},
			models.RumahSakit{Nama: "Rs. Syaiful Anwar", Start: "01:00AM", Stop: "11:00PM", Lower: "1", Upper: "5"},
			models.RumahSakit{Nama: "Rs. Panti Waluyo", Start: "01:00AM", Stop: "11:00PM", Lower: "1", Upper: "5"},
		}

		for _, rs := range RumahSakits {
			if err := seeder.Create(&rs).Error; err != nil {
				debugLog, _ := json.Marshal(err)
				log.Printf(string(debugLog))
			} else {
				log.Printf("created")
			}
		}

		var pasiens []models.Pasien = []models.Pasien{
			models.Pasien{Nama: "favian", NoHp: "08123456781", TTL: "01/02/1978", JK: "laki-laki", Kode: "1111111111", Status: "OTG", RumahSakitID: 1, AdminID: 1},
			models.Pasien{Nama: "arafat", NoHp: "08123456782", TTL: "02/03/1979", JK: "laki-laki", Kode: "2222222222", Status: "OTG", RumahSakitID: 1, AdminID: 1},
			models.Pasien{Nama: "hasan", NoHp: "08123456783", TTL: "03/04/1980", JK: "laki-laki", Kode: "3333333333", Status: "OTG", RumahSakitID: 2, AdminID: 1},
			models.Pasien{Nama: "fikri", NoHp: "08123456784", TTL: "04/05/1981", JK: "laki-laki", Kode: "4444444444", Status: "OTG", RumahSakitID: 2, AdminID: 1},
		}

		for _, pasien := range pasiens {
			if err := seeder.Create(&pasien).Error; err != nil {
				debugLog, _ := json.Marshal(err)
				log.Printf(string(debugLog))
			} else {
				log.Printf("created")
			}
		}

		var reports []models.Report = []models.Report{
			models.Report{Kode: "1111111111", RumahSakitID: 1, Longitude: "-7.9693016677787885", Latitude: "112.6356524673932", Kondisi: "sehat", Suhu: "36.5", Demam: "tidak"},
			models.Report{Kode: "2222222222", RumahSakitID: 1, Longitude: "-7.9693016677787885", Latitude: "112.6356524673932", Kondisi: "sehat", Suhu: "36.5", Demam: "tidak"},
			models.Report{Kode: "3333333333", RumahSakitID: 2, Longitude: "-7.9865676677787885", Latitude: "112.6249675673932", Kondisi: "sehat", Suhu: "36.5", Demam: "tidak"},
			models.Report{Kode: "4444444444", RumahSakitID: 2, Longitude: "-7.9865676677787885", Latitude: "112.6249675673932", Kondisi: "sehat", Suhu: "36.5", Demam: "tidak"},
		}

		for _, report := range reports {
			if err := seeder.Create(&report).Error; err != nil {
				debugLog, _ := json.Marshal(err)
				log.Printf(string(debugLog))
			} else {
				log.Printf("created")
			}
		}

		if err := seeder.Commit().Error; err != nil {
			debugLog, _ := json.Marshal(err)
			log.Printf(string(debugLog))
		} else {
			log.Printf("------------- SEED Commited ---------------")
		}
		log.Printf("------------- SEED END ---------------")
	}
}

func Truncate() {
	if bootstrap.App.ENV == "dev" || bootstrap.App.ENV == "test" {
		log.Printf("------------- TRUNCATE TABLE ---------------")
		bootstrap.App.DB.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE;")
	}
}
