package model

import "time"

type User struct {
	ID         int64      `db:"id"`
	Nickname   string     `db:"nickname"`
	Account    string     `db:"account"`
	Password   string     `db:"password"`
	Status     int        `db:"status"`
	CreateTime *time.Time `db:"createTime"`
}
