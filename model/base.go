package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

var orm *xorm.Engine
