package models

import (
	"time"
)

type (
	RumahSakit struct {
		BaseModel
		Nama         string    `json:"nama"`
		Lower        string    `json:"lower"`
		Upper        string    `json:"upper"`
		Start        string    `json:"start"`
		Stop         string    `json:"stop"`
		NextSchedule time.Time `json:"next_schedule"`
	}

	RumahSakitPaginationResponse struct {
		Meta PaginationResponse `json:"meta"`
		Data []RumahSakit       `json:"data"`
	}

	// just use string type, since it will be use on query at DB layer
	RumahSakitFilterable struct {
		Nama string `json:"nama"`
	}
)

// Callback before update RumahSakit
func (m *RumahSakit) BeforeUpdate() (err error) {
	m.UpdatedAt = time.Now()
	return
}

// Callback before create RumahSakit
func (m *RumahSakit) BeforeCreate() (err error) {
	m.CreatedAt = time.Now()
	return
}

// Create
func (m *RumahSakit) Create() (*RumahSakit, error) {
	var err error
	err = Create(&m)
	return m, err
}

// Update
func (m *RumahSakit) Update() error {
	var err error
	err = Save(&m)
	return err
}

// Delete
func (m *RumahSakit) Delete() error {
	var err error
	err = Delete(&m)
	return err
}

// FindByID
func (m *RumahSakit) FindByID(id int) (*RumahSakit, error) {
	var (
		err error
	)
	err = FindOneByID(&m, id)
	return m, err
}

// GetList
func (m *RumahSakit) GetList(page int, rp int, orderby string, sort string, filters interface{}) (interface{}, error) {
	var (
		rumahsakits []RumahSakit
		err         error
	)

	resp, err := FindAllWithPage(&rumahsakits, page, rp, orderby, sort, filters)
	return resp, err
}
