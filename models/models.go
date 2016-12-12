package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "homestead:secret@tcp(192.168.10.10:3306)/auto?charset=utf8", 30)
	orm.RunSyncdb("default", false, true) // create table
	orm.Debug = false
}
