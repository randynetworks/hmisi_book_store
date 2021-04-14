package models

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	IDCategory int
	Name       string
	Author     string
	Synopsis   string `sql:"type:text;"`
	Pages 		int
}
