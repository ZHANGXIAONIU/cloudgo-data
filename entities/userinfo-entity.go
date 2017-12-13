package entities

import (
	"time"
)

// UserInfo .
type UserInfo struct {
	UID        int        `xorm:"'uid' INT(10) PK NOT NULL AUTOINCR"`
	UserName   string     `xorm:"'username' VARCHAR(64) NULL DEFAULT NULL"`
	DepartName string     `xorm:"'departname' VARCHAR(64) NULL DEFAULT NULL"`
	CreateAt   *time.Time `xorm:"'created' DATE NULL DEFAULT NULL"`
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("UserName shold not null!")
	}
	if u.CreateAt == nil {
		t := time.Now()
		u.CreateAt = &t
	}
	return &u
}
