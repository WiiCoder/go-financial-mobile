package db

import "gorm.io/gorm"

// Role 角色
type Role struct {
	gorm.Model
	name   string // 角色名称
	remark string // 备注
}

// Menu 菜单
type Menu struct {
	gorm.Model
	ParentId int    // 父级ID
	Name     string // 菜单名称
	Type     int    // 类型： 1 - 菜单、2 - 按钮
	Sequence int    // 排序，越小越前面
}

// RoleMenu 角色菜单
type RoleMenu struct {
	gorm.Model
	RoleId int // 角色ID
	MenuId int // 菜单的ID
}

// UserRole 用户角色
type UserRole struct {
	gorm.Model
	UserId int // 系统用户ID
	RoleId int // 角色ID
}
