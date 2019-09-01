package models

import "github.com/jinzhu/gorm"

type Channel struct {
	gorm.Model
	ChannelName string
	ConnectKey  string
}
