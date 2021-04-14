package models

import "gorm.io/gorm"

//id_buku int
//id_visitor int

type Borrows struct {
	gorm.Model
	IDBook int
	IDUser int
}