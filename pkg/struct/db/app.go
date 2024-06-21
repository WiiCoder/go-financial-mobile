package db

import (
	"gorm.io/gorm"
	"time"
)

type AppUser struct {
	gorm.Model
	Username      string
	Password      string
	Salt          string
	WechatOpenid  string
	Name          string
	PhoneNumber   string
	LastLoginTime *time.Time
}
