package handler

import (
	"goconf/core/sdk/pkg"
	"gorm.io/gorm"
)

type Login struct {
	Username string `json:"username" from:"username" binding:"required"`
	Password string `json:"password" from:"password" binding:"required"`
	Code     string `json:"code" from:"code" binding:"required"`
	UUID     string `json:"uuid" from:"uuid" binding:"required"`
}

func (u *Login) GetUser(tx *gorm.DB) (user SysUser, role SysRole, err error) {
	err = tx.Table("sys_user").Where("username =? and status = '2'", u.Username).First(&user).Error
	if err != nil {
		return
	}
	_, err = pkg.CompareHashAndPassword(user.Password, u.Password)
	if err != nil {
		return
	}
	err = tx.Table("sys_role").Where("role_id =?", user.RoleId).First(&role).Error
	if err != nil {
		return
	}
	return

}
