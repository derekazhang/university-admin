package admin_user

import (
	"gorm.io/gorm"
)

type AdminUser struct {
	gorm.Model

	Account string `json:"account"`
	Pwd     string `json:"pwd,omitempty"`
	Name    string `json:"name" gorm:"type:varchar(50);not null; comment:'用户名'"`
	Phone   string `json:"phone"`
	Mail    string `json:"mail"`
	State   int    `json:"state"` // 0=冻结 1=正常
	RoleId  int    `json:"role_id" gorm:"type:int;not null; comment:'角色id'"`
}

type AdminUserFull struct {
	Token string `json:"token"`
	AdminUser
}

type AdminUserRole struct {
	gorm.Model
	Name   string `json:"name" gorm:"type:VARCHAR(30);not null;comment:'角色名称'"`
	Rules  string `json:"rules" gorm:"type:VARCHAR(255);not null;comment:'角色拥有的权限节点'"`
	Status int    `json:"status" gorm:"type:TINYINT(1);not null;default:1;comment:'状态：1 启用 0 禁用'"`
}

type AdminUserPermission struct {
	ID          int `json:"id"`          // 权限id
	Name        int `json:"name"`        // 权限名称
	Description int `json:"description"` //权限信息
}
