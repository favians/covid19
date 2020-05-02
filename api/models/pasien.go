package models

import (
	"time"

	"github.com/favians/golang_starter/bootstrap"
)

type (
	Pasien struct {
		BaseModel
		Nama         string `json:"nama"`
		NoHp         string `json:"no_hp"`
		TTL          string `json:"ttl"`
		JK           string `json:"jk"`
		Kode         string `json:"kode"`
		Status       string `json:"status"`
		RumahSakitID int    `json:"rumah_sakit_id"`
		AdminID      uint64 `json:"admin_id"`
		Email        string `json:"email"`
		Longitude    string `json:"longitude"`
		Latitude     string `json:"latitude"`
	}

	PasienPaginationResponse struct {
		Meta PaginationResponse `json:"meta"`
		Data []Pasien           `json:"data"`
	}

	// just use string type, since it will be use on query at DB layer
	PasienFilterable struct {
		Nama   string `json:"nama"`
		JK     string `json:"jk"`
		Kode   string `json:"kode"`
		Status string `json:"status"`
	}
)

// Callback before update Pasien
func (m *Pasien) BeforeUpdate() (err error) {
	m.UpdatedAt = time.Now()
	return
}

// Callback before create Pasien
func (m *Pasien) BeforeCreate() (err error) {
	m.CreatedAt = time.Now()
	return
}

// Create
func (m *Pasien) Create() (*Pasien, error) {
	var err error
	err = Create(&m)
	return m, err
}

// Update
func (m *Pasien) Update() error {
	var err error
	err = Save(&m)
	return err
}

// Delete
func (m *Pasien) Delete() error {
	var err error
	err = Delete(&m)
	return err
}

// FindByID
func (m *Pasien) FindByID(id int) (*Pasien, error) {
	var (
		err error
	)
	err = FindOneByID(&m, id)
	return m, err
}

// GetList
func (m *Pasien) GetList(page int, rp int, orderby string, sort string, filters interface{}) (interface{}, error) {
	var (
		pasiens []Pasien
		err     error
	)

	resp, err := FindAllWithPage(&pasiens, page, rp, orderby, sort, filters)
	return resp, err
}

// FindByID
func (m *Pasien) FindByCode(code string) (*Pasien, error) {
	var (
		err error
	)

	bootstrap.App.DB.Table("pasiens").Where("kode = ?", code).First(&m)
	return m, err
}
