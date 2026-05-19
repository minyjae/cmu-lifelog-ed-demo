package models

import (
	"time"

	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
)

// User table
type Users struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	Name               string    `gorm:"not null" json:"cmuitaccount_name"`
	Email              string    `gorm:"unique;not null" json:"cmuitaccount"`
	Password           string    `gorm:"" json:"-"`
	Role               string    `gorm:"default:user; not null" json:"role"`
	PreNameID          string    `gorm:"not null" json:"prename_id"`
	PreNameTH          string    `gorm:"not null" json:"prename_th"`
	PreNameEN          string    `gorm:"not null" json:"prename_en"`
	FirstNameTH        string    `gorm:"not null" json:"firstname_th"`
	FirstNameEN        string    `gorm:"not null" json:"firstname_en"`
	LastNameTH         string    `gorm:"not null" json:"lastname_th"`
	LastNameEN         string    `gorm:"not null" json:"lastname_en"`
	OrganizationCode   string    `gorm:"not null" json:"organization_code"`
	OrganizationNameTH string    `gorm:"not null" json:"organization_name_th"`
	OrganizationNameEN string    `gorm:"not null" json:"organization_name_en"`
	ITAccountTypeID    string    `gorm:"not null" json:"itaccounttype_id"`
	ITAccountTypeTH    string    `gorm:"not null" json:"itaccounttype_th"`
	ITAccountTypeEN    string    `gorm:"not null" json:"itaccounttype_en"`
	IsFirstTime        bool      `gorm:"not null;default:true" json:"is_first_time"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

func (u *Users) ToEntity() *entities.Users {
	return &entities.Users{
		ID:                 u.ID,
		Name:               u.Name,
		Email:              u.Email,
		Password:           u.Password,
		Role:               u.Role,
		PreNameID:          u.PreNameID,
		PreNameTH:          u.PreNameTH,
		PreNameEN:          u.PreNameEN,
		FirstNameTH:        u.FirstNameTH,
		FirstNameEN:        u.FirstNameEN,
		LastNameTH:         u.LastNameTH,
		LastNameEN:         u.LastNameEN,
		OrganizationCode:   u.OrganizationCode,
		OrganizationNameTH: u.OrganizationNameTH,
		OrganizationNameEN: u.OrganizationNameEN,
		ITAccountTypeID:    u.ITAccountTypeID,
		ITAccountTypeTH:    u.ITAccountTypeTH,
		ITAccountTypeEN:    u.ITAccountTypeEN,
		IsFirstTime:        u.IsFirstTime,
	}
}

func (u *Users) FromEntity(entity *entities.Users) {
	u.ID = entity.ID
	u.Name = entity.Name
	u.Email = entity.Email
	u.Password = entity.Password
	u.Role = entity.Role
	u.PreNameID = entity.PreNameID
	u.PreNameTH = entity.PreNameTH
	u.PreNameEN = entity.PreNameEN
	u.FirstNameTH = entity.FirstNameTH
	u.FirstNameEN = entity.FirstNameEN
	u.LastNameTH = entity.LastNameTH
	u.LastNameEN = entity.LastNameEN
	u.OrganizationCode = entity.OrganizationCode
	u.OrganizationNameTH = entity.OrganizationNameTH
	u.OrganizationNameEN = entity.OrganizationNameEN
	u.ITAccountTypeID = entity.ITAccountTypeID
	u.ITAccountTypeTH = entity.ITAccountTypeTH
	u.ITAccountTypeEN = entity.ITAccountTypeEN
	u.IsFirstTime = entity.IsFirstTime
}
