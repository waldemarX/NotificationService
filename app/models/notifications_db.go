package model

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model

	ID       int `gorm:"primarykey"`
	Type 	 string
	To 		 int
	From 	 int
}
