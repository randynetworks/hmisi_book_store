package models

import "gorm.io/gorm"

//sosial_id
//name string 200
//email string 200
//No_telp string 200
//address text
//role bool default = 0

type Users struct {
	gorm.Model
	SocialID string
	Name 	 string
	Email 	 string
	NoTelp 	 string
	Address  string `sql:"type:text;"`
	Role 	 bool	`gorm:"default:0"`
}


