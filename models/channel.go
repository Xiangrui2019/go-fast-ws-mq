package models

import (
	"github.com/jinzhu/gorm"
)

type Channel struct {
	gorm.Model
	Name        string
	Description string
	ConnectKey  string
}
