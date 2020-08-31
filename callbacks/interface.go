package callbacks

import "github.com/nhokboo/gorm"

type beforeSaveInterface interface {
	BeforeSave(*gorm.DB) error
}

type beforeCreateInterface interface {
	BeforeCreate(*gorm.DB) error
}
