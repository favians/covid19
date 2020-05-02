package models

import (
	"time"
)

type (
	Admin struct {
		BaseModel
		Name         string `json:"name"`
		Username     string `json:"username"`
		Password     string `json:"password" valid:"password"`
		RumahSakitID int    `json:"rumah_sakit_id"`
	}

	AdminPaginationResponse struct {
		Meta PaginationResponse `json:"meta"`
		Data []Admin            `json:"data"`
	}

	// just use string type, since it will be use on query at DB layer
	AdminFilterable struct {
		Name         string `json:"name"`
		RumahSakitID string `json:"rumah_sakit_id"`
	}
)

// Callback before update Admin
func (m *Admin) BeforeUpdate() (err error) {
	m.UpdatedAt = time.Now()
	return
}

// Callback before create Admin
func (m *Admin) BeforeCreate() (err error) {
	m.CreatedAt = time.Now()
	return
}

// Create
func (m *Admin) Create() (*Admin, error) {
	var err error
	err = Create(&m)
	return m, err
}

// Update
func (m *Admin) Update() error {
	var err error
	err = Save(&m)
	return err
}

// Delete
func (m *Admin) Delete() error {
	var err error
	err = Delete(&m)
	return err
}

// FindByID
func (m *Admin) FindByID(id int) (*Admin, error) {
	var (
		err error
	)
	err = FindOneByID(&m, id)
	return m, err
}

// GetList
func (m *Admin) GetList(page int, rp int, orderby string, sort string, filters interface{}) (interface{}, error) {
	var (
		admins []Admin
		err    error
	)

	resp, err := FindAllWithPage(&admins, page, rp, orderby, sort, filters)
	return resp, err
}
