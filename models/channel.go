package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Channel struct {
	gorm.Model
	Name        string
	Description string
	ConnectKey  string
	CreateTime  time.Duration
}
