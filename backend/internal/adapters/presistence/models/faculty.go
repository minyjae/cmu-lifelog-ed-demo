package models

import (
	"github.com/minyjae/cmu-life-long-ed-api/internal/core/domain/entities"
)

type Faculty struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Code   string `gorm:"not null" json:"code"`
	NameTH string `gorm:"not null" json:"nameTH"`
	NameEN string `gorm:"not null" json:"nameEN"`
}

func (f *Faculty) ToEntity() *entities.Faculty {
	return &entities.Faculty{
		ID:     f.ID,
		Code:   f.Code,
		NameTH: f.NameTH,
		NameEN: f.NameEN,
	}
}

func (f *Faculty) FromEntity(entity *entities.Faculty) {
	f.ID = entity.ID
	f.Code = entity.Code
	f.NameTH = entity.NameTH
	f.NameEN = entity.NameEN
}
