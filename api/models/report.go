package models

import (
	"time"
)

type (
	Report struct {
		BaseModel
		Nama   string `json:"nama"`
		NoHp   string `json:"no_hp"`
		TTL    string `json:"ttl"`
		JK     string `json:"jk"`
		Kode   string `json:"kode"`
		Status string `json:"status"`
	}

	ReportPaginationResponse struct {
		Meta PaginationResponse `json:"meta"`
		Data []Report           `json:"data"`
	}

	// just use string type, since it will be use on query at DB layer
	ReportFilterable struct {
		Nama   string `json:"nama"`
		JK     string `json:"jk"`
		Kode   string `json:"kode"`
		Status string `json:"status"`
	}
)

// Callback before update Report
func (m *Report) BeforeUpdate() (err error) {
	m.UpdatedAt = time.Now()
	return
}

// Callback before create Report
func (m *Report) BeforeCreate() (err error) {
	m.CreatedAt = time.Now()
	return
}

// Create
func (m *Report) Create() (*Report, error) {
	var err error
	err = Create(&m)
	return m, err
}

// Update
func (m *Report) Update() error {
	var err error
	err = Save(&m)
	return err
}

// Delete
func (m *Report) Delete() error {
	var err error
	err = Delete(&m)
	return err
}

// FindByID
func (m *Report) FindByID(id int) (*Report, error) {
	var (
		err error
	)
	err = FindOneByID(&m, id)
	return m, err
}

// GetList
func (m *Report) GetList(page int, rp int, orderby string, sort string, filters interface{}) (interface{}, error) {
	var (
		reports []Report
		err     error
	)

	resp, err := FindAllWithPage(&reports, page, rp, orderby, sort, filters)
	return resp, err
}
