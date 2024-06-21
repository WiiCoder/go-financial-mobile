package db

import (
	"gorm.io/gorm"
	"time"
)

// Distributor 渠道
type Distributor struct {
	gorm.Model
	Name string // 渠道名称
}

// Partner 股东
type Partner struct {
	gorm.Model
	name string // 股东名称
}

// SystemConfig 系统配置
type SystemConfig struct {
	gorm.Model
	Name   string // 名称
	Key    string // 键
	Value  string // 值
	Remark string // 备注
}

// SystemUser 系统用户
type SystemUser struct {
	gorm.Model
	Username      string    // 用户名
	Password      string    // 密码
	Salt          string    // 盐值
	LastLoginTime time.Time // 最近一次登陆时间
}

// UserStore 用户门店
type UserStore struct {
	gorm.Model
	UserId  int // 系统用户ID
	StoreId int // 门店ID
}
