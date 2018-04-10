package model

import (
	"time"
)

//用户

type User struct {
	ID         int64      `form:"userId" xorm:"pk not null autoincr"`
	Nickname   string     `form:"nickname" xorm:"varchar(20) not null"`
	Account    string     `form:"account" xorm:"varchar(60) not null"`
	Password   string     `form:"password" xorm:"varchar(60) not null"`
	Status     int        `form:"status" xorm:"int(1) default 0"`
	CreateTime *time.Time `xorm:"datetime created"`
}

//查询用户

func (u *User) GetUser(id int64) (*User, error) {
	user := &User{ID: id}
	_, err := orm.Get(user)
	return user, err
}

//添加用户

// func (u *User) InsertUser() error {

// }

//
